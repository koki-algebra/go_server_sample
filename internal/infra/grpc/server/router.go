package server

import (
	"context"
	"database/sql"
	"net/http"

	userv1 "github.com/koki-algebra/go_server_sample/internal/infra/grpc/generated/user/v1/userv1connect"
	"github.com/koki-algebra/go_server_sample/internal/infra/grpc/service"
	"github.com/koki-algebra/go_server_sample/internal/infra/repository"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
)

func newRouter(ctx context.Context, sqldb *sql.DB) http.Handler {
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

	return mux
}
