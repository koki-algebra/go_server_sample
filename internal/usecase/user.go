package usecase

import (
	"context"

	"github.com/google/uuid"

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

type SaveUserInput struct {
	ID   *uuid.UUID
	Name *string
}

func (u *User) Save(ctx context.Context, input SaveUserInput) (*entity.User, error) {
	user := new(entity.User)
	if input.ID != nil {
		user.ID = *input.ID
	}
	if input.Name != nil {
		user.Name = *input.Name
	}

	if err := u.repo.Save(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
