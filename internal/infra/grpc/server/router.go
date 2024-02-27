package server

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/httplog/v2"

	userv1 "github.com/koki-algebra/go_server_sample/internal/infra/grpc/generated/user/v1/userv1connect"
	"github.com/koki-algebra/go_server_sample/internal/infra/grpc/service"
	"github.com/koki-algebra/go_server_sample/internal/infra/middleware"
	"github.com/koki-algebra/go_server_sample/internal/infra/repository"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
)

func newRouter(ctx context.Context, sqldb *sql.DB) http.Handler {
	logger := httplog.NewLogger("graphql", httplog.Options{
		LogLevel:         slog.LevelInfo,
		LevelFieldName:   "severity",
		MessageFieldName: "message",
		JSON:             true,
		Concise:          false,
		RequestHeaders:   true,
		TimeFieldFormat:  time.RFC3339,
		TimeFieldName:    "time",
		QuietDownRoutes: []string{
			"/",
			"/ping",
		},
		QuietDownPeriod: 10 * time.Second,
		SourceFieldName: "logging.googleapis.com/sourceLocation",
	})

	var (
		// repository
		userRepository = repository.NewUserRepository(sqldb)

		// usecase
		user = usecase.NewUser(userRepository)

		// services
		userService = service.NewUserService(user)
	)

	mux := http.NewServeMux()
	mux.Handle(userv1.NewUserServiceHandler(userService))

	return middleware.With(mux, httplog.RequestLogger(logger))
}
