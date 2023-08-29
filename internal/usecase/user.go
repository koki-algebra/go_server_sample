package usecase

import (
	"context"

	"github.com/koki-algebra/grpc_sample/internal/entity"
	"github.com/koki-algebra/grpc_sample/internal/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{
		repo: repository.NewUserRepository(),
	}
}

func (u *UserUsecase) GetByID(ctx context.Context, id string) (*entity.User, error) {
	user, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
