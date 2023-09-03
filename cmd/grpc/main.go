package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/koki-algebra/go_server_sample/internal/infra/grpc/server"
)

func main() {
	// logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx := context.Background()

	// start server
	srv := server.NewServer(8080)
	if err := srv.Run(ctx); err != nil {
		logger.Error(err.Error())
	}
}
