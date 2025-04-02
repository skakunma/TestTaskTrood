package storage

import (
	"context"
	"database/sql"
)

type (
	Storage interface {
		GetAnswer(ctx context.Context, question string) (string, error)
	}
	PostgresStorage struct {
		Db *sql.DB
	}
)
