package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func Open(ctx context.Context, host, port, user, password, database string) (*sql.DB, error) {
	// connect to database
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, database)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// verify connection
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func OpenBun(db *sql.DB) *bun.DB {
	return bun.NewDB(db, pgdialect.New())
}
