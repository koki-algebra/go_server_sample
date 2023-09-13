package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/koki-algebra/go_server_sample/internal/infra/database"
	groupv1 "github.com/koki-algebra/go_server_sample/internal/infra/grpc/generated/group/v1/v1connect"
	userv1 "github.com/koki-algebra/go_server_sample/internal/infra/grpc/generated/user/v1/v1connect"
	"github.com/koki-algebra/go_server_sample/internal/infra/grpc/service"
	"github.com/koki-algebra/go_server_sample/internal/infra/repository"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
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
	mux := http.NewServeMux()

	// database
	db, err := database.Open(
		ctx,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
	if err != nil {
		return err
	}
	defer db.Close()

	// repository
	userRepository := repository.NewUserRepository(db)
	groupRepository := repository.NewGroupRepository(db)

	// usecases
	user := usecase.NewUser(userRepository)
	group := usecase.NewGroup(groupRepository)

	// services
	userService := service.NewUserService(user)
	groupService := service.NewGroupService(group)

	// handlers
	mux.Handle(userv1.NewUserServiceHandler(userService))
	mux.Handle(groupv1.NewGroupServiceHandler(groupService))

	srv := &http.Server{
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		Addr:              fmt.Sprintf(":%d", s.port),
		WriteTimeout:      time.Second * 60,
		ReadTimeout:       time.Second * 15,
		ReadHeaderTimeout: time.Second * 15,
		IdleTimeout:       time.Second * 120,
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		slog.Info(fmt.Sprintf("start Connect server port: %d", s.port))
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

	return eg.Wait()
}
