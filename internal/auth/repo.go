package auth

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
}

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repo{db: db}
}
