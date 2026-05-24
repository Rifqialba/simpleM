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
			display_name,
			password_hash
		)
		VALUES ($1, $2, $3, $4)
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
		user.PasswordHash,
	).Scan(
		&user.ID,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return err
}

func (r *Repository) FindByEmail(
	ctx context.Context,
	email string,
) (*User, error) {

	query := `
		SELECT
			id,
			email,
			username,
			display_name,
			avatar_url,
			password_hash,
			is_active,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
	`

	user := &User{}

	err := r.db.QueryRow(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.DisplayName,
		&user.AvatarURL,
		&user.PasswordHash,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}