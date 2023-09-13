package repository

import (
	"context"

	"github.com/koki-algebra/go_server_sample/internal/entity"
)

type GroupRepository interface {
	FindByID(ctx context.Context, id string) (*entity.Group, error)
}
