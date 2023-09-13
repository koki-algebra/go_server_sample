package repository

import (
	"context"

	"github.com/koki-algebra/go_server_sample/internal/entity"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*entity.User, error)
}
