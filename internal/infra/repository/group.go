package repository

import (
	"context"

	"github.com/koki-algebra/go_server_sample/internal/entity"
	model "github.com/koki-algebra/go_server_sample/internal/infra/database/generated/sqlboiler"
	"github.com/koki-algebra/go_server_sample/internal/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func NewGroupRepository(exec boil.ContextExecutor) repository.GroupRepository {
	return &groupRepositoryImpl{
		exec: exec,
	}
}

type groupRepositoryImpl struct {
	exec boil.ContextExecutor
}

func (r *groupRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.Group, error) {
	group, err := model.Groups(qm.Where("id = ?", id)).One(ctx, r.exec)
	if err != nil {
		return nil, err
	}

	return convertGroup(group), nil
}

func convertGroup(group *model.Group) *entity.Group {
	if group == nil {
		return nil
	}

	return &entity.Group{
		ID:   group.ID,
		Name: group.Name,
	}
}
