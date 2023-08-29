package service

import (
	"context"

	"github.com/koki-algebra/grpc_sample/internal/infra/grpc/generated"
	"github.com/koki-algebra/grpc_sample/internal/usecase"
)

type UserService struct {
	usecase *usecase.UserUsecase
	generated.UnimplementedUserServiceServer
}

func NewUserService(usecase *usecase.UserUsecase) *UserService {
	return &UserService{
		usecase: usecase,
	}
}

func (s *UserService) GetByID(ctx context.Context, req *generated.GetByIDRequest) (*generated.GetByIDResponse, error) {
	return &generated.GetByIDResponse{
		Id:   "xyz",
		Name: "foo",
	}, nil
}
