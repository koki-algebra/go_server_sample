package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/sourcegraph/conc/pool"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/koki-algebra/go_server_sample/internal/infra/config"
	"github.com/koki-algebra/go_server_sample/internal/infra/database"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s Server) Run(ctx context.Context) error {
	// database
	db, err := database.Open(ctx)
	if err != nil {
		return err
	}
	defer db.Close()

	router := newRouter(ctx, db)

	srv := &http.Server{
		Handler:           h2c.NewHandler(router, &http2.Server{}),
		Addr:              fmt.Sprintf(":%d", config.Env.ServerPort),
		WriteTimeout:      time.Second * 60,
		ReadTimeout:       time.Second * 15,
		ReadHeaderTimeout: time.Second * 15,
		IdleTimeout:       time.Second * 120,
	}

	pool := pool.New().WithErrors().WithContext(ctx)
	pool.Go(func(ctx context.Context) error {
		slog.InfoContext(ctx, fmt.Sprintf("start HTTP server port: %d", config.Env.ServerPort))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	<-ctx.Done()
	slog.Info("stopping Connect server...")
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown", "error", err)
	}

	return pool.Wait()
}
