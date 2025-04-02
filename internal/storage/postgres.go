package storage

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgresStorage(dbURL string) (Storage, error) {
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return nil, err
	}
	storage := &PostgresStorage{Db: db}
	err = storage.createSchema()
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *PostgresStorage) createSchema() error {
	query := `
	CREATE TABLE IF NOT EXISTS answers (
	question TEXT,
	answer TEXT);
	`
	_, err := s.Db.Query(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStorage) GetAnswer(ctx context.Context, question string) (string, error) {
	var ancwer string
	err := s.Db.QueryRowContext(ctx, "SELECT answer FROM answers WHERE question LIKE $1", question).Scan(&ancwer)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return "", nil
		}
		return "", err
	}
	return ancwer, nil
}
