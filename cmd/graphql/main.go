package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/koki-algebra/go_server_sample/internal/infra/config"
	"github.com/koki-algebra/go_server_sample/internal/infra/graphql/server"
)

func main() {
	if err := run(context.Background()); err != nil {
		slog.Error("failed to terminated server", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	if err := config.Init(); err != nil {
		return err
	}

	srv := server.NewServer()
	return srv.Run(ctx)
}
