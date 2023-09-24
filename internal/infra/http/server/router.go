package server

import (
	"database/sql"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/koki-algebra/go_server_sample/internal/infra/http/controller"
	"github.com/koki-algebra/go_server_sample/internal/infra/http/oapi"
	"github.com/koki-algebra/go_server_sample/internal/infra/repository"
	"github.com/koki-algebra/go_server_sample/internal/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
)

func newRouter(db *sql.DB) (*chi.Mux, error) {
	r := chi.NewRouter()

	swagger, err := oapi.GetSwagger()
	if err != nil {
		return nil, err
	}
	swagger.Servers = nil
	r.Use(middleware.OapiRequestValidator(swagger))

	// logger
	logger := httplog.NewLogger("app", httplog.Options{
		JSON: true,
	})
	r.Use(httplog.RequestLogger(logger))

	// user handler
	user := usecase.NewUser(repository.NewUserRepository(db))

	ctrl := controller.New(user)
	oapi.HandlerFromMux(ctrl, r)

	return r, nil
}
