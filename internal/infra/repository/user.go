package repository

import (
	"context"
	"database/sql"

	"github.com/koki-algebra/go_server_sample/internal/entity"
	"github.com/koki-algebra/go_server_sample/internal/repository"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func NewUserRepository(sqldb *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{
		db: bun.NewDB(sqldb, pgdialect.New()),
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
