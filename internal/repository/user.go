package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/koki-algebra/go_server_sample/internal/entity"
)

type UserRepository interface {
	Save(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
}
