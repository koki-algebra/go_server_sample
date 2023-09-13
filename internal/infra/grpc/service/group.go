package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/koki-algebra/go_server_sample/internal/entity"
	groupv1 "github.com/koki-algebra/go_server_sample/internal/infra/grpc/generated/group/v1"
	"github.com/koki-algebra/go_server_sample/internal/infra/grpc/generated/group/v1/v1connect"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
)

type GroupService struct {
	usecase *usecase.Group
}

func NewGroupService(usecase *usecase.Group) v1connect.GroupServiceHandler {
	return &GroupService{
		usecase: usecase,
	}
}

func (s *GroupService) GetByID(ctx context.Context, req *connect.Request[groupv1.GetByIDRequest]) (*connect.Response[groupv1.GetByIDResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	group, err := s.usecase.GetByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return convertGroup(group), nil
}

func convertGroup(group *entity.Group) *connect.Response[groupv1.GetByIDResponse] {
	if group == nil {
		return nil
	}

	return connect.NewResponse[groupv1.GetByIDResponse](&groupv1.GetByIDResponse{
		Id:   group.ID,
		Name: group.Name,
	})
}
