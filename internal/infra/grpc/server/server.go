package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	userpb "github.com/koki-algebra/go_server_sample/internal/infra/grpc/generated/user/v1"
	"github.com/koki-algebra/go_server_sample/internal/infra/grpc/service"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	srv := grpc.NewServer()

	// usecases
	user := usecase.NewUser()

	// services
	userService := service.NewUserService(user)

	// register services
	userpb.RegisterUserServiceServer(srv, userService)

	reflection.Register(srv)

	go func() {
		slog.Info(fmt.Sprintf("start gRPC server port: %d", s.port))
		srv.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("stopping gRPC server...")
	srv.GracefulStop()

	return nil
}
