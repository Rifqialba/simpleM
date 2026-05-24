package room

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
	room *Room,
) error {

	query := `
		INSERT INTO rooms (
			workspace_id,
			created_by,
			name,
			description
		)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id,
			is_archived,
			created_at,
			updated_at
	`

	return r.db.QueryRow(
		ctx,
		query,
		room.WorkspaceID,
		room.CreatedBy,
		room.Name,
		room.Description,
	).Scan(
		&room.ID,
		&room.IsArchived,
		&room.CreatedAt,
		&room.UpdatedAt,
	)
}