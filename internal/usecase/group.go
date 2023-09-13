package usecase

import (
	"context"

	"github.com/koki-algebra/go_server_sample/internal/entity"
	"github.com/koki-algebra/go_server_sample/internal/repository"
)

type Group struct {
	repo repository.GroupRepository
}

func NewGroup(repo repository.GroupRepository) *Group {
	return &Group{
		repo: repo,
	}
}

func (g *Group) GetByID(ctx context.Context, id string) (*entity.Group, error) {
	group, err := g.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return group, nil
}
