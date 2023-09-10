package repository

import (
	"context"

	"github.com/koki-algebra/go_server_sample/internal/entity"
	"github.com/uptrace/bun"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*entity.User, error)
}

func NewUserRepository(db *bun.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

type userRepositoryImpl struct {
	db *bun.DB
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	if err := r.db.
		NewSelect().
		Model(&user).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, err
	}

	return &user, nil
}
