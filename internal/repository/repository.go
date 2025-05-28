package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ochko-b/goapp/generated/sqlc"
)

type Repository struct {
	*sqlc.Queries
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		Queries: sqlc.New(db),
		db:      db,
	}
}

func (r *Repository) WithTx(tx pgx.Tx) *Repository {
	return &Repository{
		Queries: r.Queries.WithTx(tx),
		db:      r.db,
	}
}

func (r *Repository) BeginTx(ctx context.Context) (pgx.Tx, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
