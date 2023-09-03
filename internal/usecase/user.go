package usecase

import (
	"context"

	"github.com/koki-algebra/grpc_sample/internal/entity"
	"github.com/koki-algebra/grpc_sample/internal/repository"
)

type User struct {
	repo repository.UserRepository
}

func NewUser() *User {
	return &User{
		repo: repository.NewUserRepository(),
	}
}

func (u *User) GetByID(ctx context.Context, id string) (*entity.User, error) {
	user, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
