package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"

	"github.com/koki-algebra/go_server_sample/internal/entity"
	userv1 "github.com/koki-algebra/go_server_sample/internal/infra/grpc/generated/user/v1"
	"github.com/koki-algebra/go_server_sample/internal/infra/grpc/generated/user/v1/userv1connect"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
)

type UserService struct {
	user *usecase.User
}

func NewUserService(user *usecase.User) userv1connect.UserServiceHandler {
	return &UserService{
		user: user,
	}
}

func (s *UserService) GetByID(
	ctx context.Context,
	req *connect.Request[userv1.GetByIDRequest],
) (*connect.Response[userv1.GetByIDResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	id, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	user, err := s.user.GetByID(ctx, id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return convertUser(user), nil
}

func convertUser(user *entity.User) *connect.Response[userv1.GetByIDResponse] {
	if user == nil {
		return nil
	}

	return connect.NewResponse[userv1.GetByIDResponse](&userv1.GetByIDResponse{
		Id:   user.ID.String(),
		Name: user.Name,
	})
}
