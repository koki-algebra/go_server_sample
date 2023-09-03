package repository

import (
	"context"

	"github.com/koki-algebra/go_server_sample/internal/entity"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*entity.User, error)
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

type userRepositoryImpl struct{}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.User, error) {
	return &entity.User{
		ID:   "id",
		Name: "name",
	}, nil
}
