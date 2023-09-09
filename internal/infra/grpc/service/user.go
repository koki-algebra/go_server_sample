package service

import (
	"context"

	"github.com/koki-algebra/go_server_sample/internal/entity"
	pb "github.com/koki-algebra/go_server_sample/internal/infra/grpc/generated/user/v1"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
)

type UserService struct {
	usecase *usecase.User
	pb.UnimplementedUserServiceServer
}

func NewUserService(usecase *usecase.User) *UserService {
	return &UserService{
		usecase: usecase,
	}
}

func (s *UserService) GetByID(ctx context.Context, req *pb.GetByIDRequest) (*pb.GetByIDResponse, error) {
	user, err := s.usecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return convertUser(user), nil
}

func convertUser(user *entity.User) *pb.GetByIDResponse {
	if user == nil {
		return nil
	}

	return &pb.GetByIDResponse{
		Id:   user.ID,
		Name: user.Name,
	}
}
