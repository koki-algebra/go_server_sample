package controller

import (
	"github.com/koki-algebra/go_server_sample/internal/infra/http/oapi"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
)

func New(user *usecase.User) oapi.ServerInterface {
	return &controllerImpl{
		user: user,
	}
}

type controllerImpl struct {
	user *usecase.User
}
