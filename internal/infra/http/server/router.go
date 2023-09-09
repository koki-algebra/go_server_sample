package server

import (
	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/koki-algebra/go_server_sample/internal/infra/http/generated"
	"github.com/koki-algebra/go_server_sample/internal/infra/http/handler"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
)

func newRouter(swagger *openapi3.T) *chi.Mux {
	router := chi.NewRouter()

	// Use our validation middleware to check all requests against the OpenAPI schema.
	router.Use(middleware.OapiRequestValidator(swagger))
	router.Use()

	// user handler
	userUsecase := usecase.NewUser()
	user := handler.NewUser(userUsecase)
	generated.HandlerFromMux(user, router)

	return router
}
