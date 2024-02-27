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
	"github.com/koki-algebra/go_server_sample/internal/infra/http/controller"
	"github.com/koki-algebra/go_server_sample/internal/infra/http/oapi"
	"github.com/koki-algebra/go_server_sample/internal/infra/repository"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
)

func newRouter(ctx context.Context, sqldb *sql.DB) (http.Handler, error) {
	// Logger
	logger := httplog.NewLogger("http", httplog.Options{
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

	// Initialize router
	r := chi.NewRouter()

	swagger, err := oapi.GetSwagger()
	if err != nil {
		return nil, err
	}
	swagger.Servers = nil

	// Apply middleware
	r.Use(
		httplog.RequestLogger(logger),
		middleware.Heartbeat("/ping"),
		cors.New(cors.Options{
			AllowedOrigins: strings.Split(config.Env.ServerAllowOrigins, ","),
			AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
			AllowedHeaders: []string{"Authorization", "Content-Type"},
		}).Handler,
	)

	// user handler
	user := usecase.NewUser(repository.NewUserRepository(sqldb))

	ctrl := controller.New(user)
	oapi.HandlerFromMux(ctrl, r)

	return r, nil
}
