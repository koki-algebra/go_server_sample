package usecase

import "github.com/koki-algebra/grpc_sample/internal/repository"

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{
		repo: repository.NewUserRepository(),
	}
}
