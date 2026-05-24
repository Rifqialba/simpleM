package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(
	ctx context.Context,
	user *User,
) error {

	query := `
		INSERT INTO users (
			email,
			username,
			display_name
		)
		VALUES ($1, $2, $3)
		RETURNING
			id,
			is_active,
			created_at,
			updated_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		user.Email,
		user.Username,
		user.DisplayName,
	).Scan(
		&user.ID,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return err
}