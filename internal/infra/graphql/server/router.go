package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/koki-algebra/go_server_sample/internal/infra/graphql/generated"
	"github.com/koki-algebra/go_server_sample/internal/infra/graphql/resolver"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
	"github.com/rs/cors"
)

func NewRouter(ctx context.Context) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("ping")); err != nil {
			slog.Error("error in writing response body", "error", fmt.Sprintf("%+v", err))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	})

	// usecases
	user := usecase.NewUser()

	// resolvers
	resolvers := resolver.New(user)

	cfg := generated.Config{Resolvers: resolvers}

	gqlSrv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	mux.Handle("/graphql", gqlSrv)
	mux.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))

	router := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	}).Handler(mux)

	return router
}
