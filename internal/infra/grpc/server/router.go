package server

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/rs/cors"

	"github.com/koki-algebra/go_server_sample/internal/infra/config"
	userv1 "github.com/koki-algebra/go_server_sample/internal/infra/grpc/generated/user/v1/v1connect"
	"github.com/koki-algebra/go_server_sample/internal/infra/grpc/service"
	"github.com/koki-algebra/go_server_sample/internal/infra/repository"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
)

func newRouter(ctx context.Context, sqldb *sql.DB) http.Handler {
	// Logger
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

	// Initialize router
	router := chi.NewRouter()

	// Apply middleware
	router.Use(
		httplog.RequestLogger(logger),
		middleware.Heartbeat("/ping"),
		cors.New(cors.Options{
			AllowedOrigins: strings.Split(config.Env.ServerAllowOrigins, ","),
			AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
			AllowedHeaders: []string{"Authorization", "Content-Type"},
		}).Handler,
	)

	// handlers
	router.Handle(userv1.NewUserServiceHandler(userService))

	return router
}
