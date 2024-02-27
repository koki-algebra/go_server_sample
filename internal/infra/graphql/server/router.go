package server

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/rs/cors"

	"github.com/koki-algebra/go_server_sample/internal/infra/config"
	"github.com/koki-algebra/go_server_sample/internal/infra/graphql/generated"
	"github.com/koki-algebra/go_server_sample/internal/infra/graphql/resolver"
	"github.com/koki-algebra/go_server_sample/internal/infra/middleware"
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
	)

	// resolvers
	resolvers := resolver.New(user)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolvers,
	}))

	// Initialize router
	mux := http.NewServeMux()

	mux.Handle("/graphql", srv)
	mux.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))

	return middleware.With(mux,
		httplog.RequestLogger(logger),
		chiMiddleware.Heartbeat("/ping"),
		cors.New(cors.Options{
			AllowedOrigins: strings.Split(config.Env.ServerAllowOrigins, ","),
			AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
			AllowedHeaders: []string{"Authorization", "Content-Type"},
		}).Handler,
	)
}
