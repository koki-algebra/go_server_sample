package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/koki-algebra/go_server_sample/internal/infra/http/server"
)

func main() {
	if err := run(context.Background()); err != nil {
		slog.Error("failed to terminated server", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	srv := server.NewServer(8080)
	return srv.Run(ctx)
}
