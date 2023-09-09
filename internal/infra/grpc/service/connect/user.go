package connect

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/koki-algebra/go_server_sample/internal/entity"
	userv1 "github.com/koki-algebra/go_server_sample/internal/infra/grpc/generated/user/v1"
	"github.com/koki-algebra/go_server_sample/internal/infra/grpc/generated/user/v1/v1connect"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
)

type UserService struct {
	usecase *usecase.User
}

func NewUserService(usecase *usecase.User) v1connect.UserServiceHandler {
	return &UserService{
		usecase: usecase,
	}
}

func (s *UserService) GetByID(ctx context.Context, req *connect.Request[userv1.GetByIDRequest]) (*connect.Response[userv1.GetByIDResponse], error) {
	user, err := s.usecase.GetByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	return convertUser(user), nil
}

func convertUser(user *entity.User) *connect.Response[userv1.GetByIDResponse] {
	if user == nil {
		return nil
	}

	return &connect.Response[userv1.GetByIDResponse]{
		Msg: &userv1.GetByIDResponse{
			Id:   user.ID,
			Name: user.Name,
		},
	}
}
