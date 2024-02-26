package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/koki-algebra/go_server_sample/internal/infra/config"

	_ "github.com/lib/pq"
)

func Open(ctx context.Context) (*sql.DB, error) {
	// connect to database
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		config.Env.DBUser,
		config.Env.DBPassword,
		config.Env.DBHost,
		config.Env.DBPort,
		config.Env.DBDatabase,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// verify connection
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
