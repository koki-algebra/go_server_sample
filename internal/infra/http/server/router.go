package server

import (
	"database/sql"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/koki-algebra/go_server_sample/internal/infra/database"
	"github.com/koki-algebra/go_server_sample/internal/infra/http/generated"
	"github.com/koki-algebra/go_server_sample/internal/infra/http/handler"
	"github.com/koki-algebra/go_server_sample/internal/infra/repository"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
)

func newRouter(swagger *openapi3.T, sqldb *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	// Use our validation middleware to check all requests against the OpenAPI schema.
	router.Use(middleware.OapiRequestValidator(swagger))
	router.Use()

	bundb := database.OpenBun(sqldb)

	// repository
	userRepository := repository.NewUserRepository(bundb)

	// user handler
	userUsecase := usecase.NewUser(userRepository)
	user := handler.NewUser(userUsecase)
	generated.HandlerFromMux(user, router)

	return router
}
