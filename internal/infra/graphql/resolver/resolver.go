package resolver

import "github.com/koki-algebra/grpc_sample/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	user *usecase.User
}

func New(user *usecase.User) *Resolver {
	return &Resolver{
		user: user,
	}
}
