package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"

	"github.com/koki-algebra/go_server_sample/internal/entity"
	"github.com/koki-algebra/go_server_sample/internal/repository"
)

func NewUserRepository(sqldb *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{
		db: bun.NewDB(sqldb, pgdialect.New()),
	}
}

type userRepositoryImpl struct {
	db *bun.DB
}

func (r *userRepositoryImpl) Save(ctx context.Context, user *entity.User) error {
	return nil
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
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
