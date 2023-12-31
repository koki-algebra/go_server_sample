package usecase

import (
	"context"

	"github.com/koki-algebra/go_server_sample/internal/entity"
	"github.com/koki-algebra/go_server_sample/internal/repository"
)

type User struct {
	repo repository.UserRepository
}

func NewUser(repo repository.UserRepository) *User {
	return &User{
		repo: repo,
	}
}

func (u *User) GetByID(ctx context.Context, id string) (*entity.User, error) {
	user, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
