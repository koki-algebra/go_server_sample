package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/koki-algebra/go_server_sample/internal/infra/database"
)

type Server struct {
	port int
}

func NewServer(port int) *Server {
	return &Server{
		port: port,
	}
}

func (s Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// connect to database
	db, err := database.Open(ctx)
	if err != nil {
		return err
	}
	defer db.Close()

	router := NewRouter(ctx, db)
	srv := http.Server{
		Addr:              fmt.Sprintf(":%d", s.port),
		WriteTimeout:      time.Second * 60,
		ReadTimeout:       time.Second * 15,
		ReadHeaderTimeout: time.Second * 15,
		IdleTimeout:       time.Second * 120,
		Handler:           router,
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		slog.Info(fmt.Sprintf("start GraphQL server port: %d", s.port))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	<-ctx.Done()
	slog.Info("stopping GraphQL server...")
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown", "error", err)
	}

	return eg.Wait()
}
