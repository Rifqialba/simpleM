package workspace

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(
	db *pgxpool.Pool,
) *Repository {

	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(
	ctx context.Context,
	workspace *Workspace,
) error {

	query := `
		INSERT INTO workspaces (
			owner_id,
			name,
			slug,
			description,
			is_personal
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING
			id,
			created_at,
			updated_at
	`

	return r.db.QueryRow(
		ctx,
		query,
		workspace.OwnerID,
		workspace.Name,
		workspace.Slug,
		workspace.Description,
		workspace.IsPersonal,
	).Scan(
		&workspace.ID,
		&workspace.CreatedAt,
		&workspace.UpdatedAt,
	)
}