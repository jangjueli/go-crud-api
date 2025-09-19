package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Create(ctx context.Context, u User) (int64, error)
	Update(ctx context.Context, u User) (*User, error)
	Delete(ctx context.Context, id int64) error
}

type Repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &Repo{db: db}
}

func (r *Repo) GetAll(ctx context.Context) ([]User, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *Repo) GetByID(ctx context.Context, id int64) (*User, error) {
	row := r.db.QueryRow(ctx, "SELECT id, name FROM users WHERE id=$1", id)

	var u User
	if err := row.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repo) Create(ctx context.Context, u User) (int64, error) {
	row := r.db.QueryRow(ctx, "INSERT INTO users (name) VALUES ($1) RETURNING id", u.Name)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repo) Update(ctx context.Context, u User) (*User, error) {
	row := r.db.QueryRow(ctx, "UPDATE users SET name=$1 WHERE id=$2 RETURNING id, name", u.Name, u.ID)

	var updated User
	if err := row.Scan(&updated.ID, &updated.Name); err != nil {
		return nil, err // row ไม่เจอ หรือ error DB
	}

	return &updated, nil
}

func (r *Repo) Delete(ctx context.Context, id int64) error {
	ct, err := r.db.Exec(ctx, "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	return nil
}
