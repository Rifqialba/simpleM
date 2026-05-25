package tab

import (
	"context"
	"encoding/json"

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

func (r *Repository) GetNextPosition(
	ctx context.Context,
	roomID string,
) (int, error) {

	query := `
		SELECT COALESCE(MAX(position), 0) + 1
		FROM room_tabs
		WHERE room_id = $1
	`

	var position int

	err := r.db.QueryRow(
		ctx,
		query,
		roomID,
	).Scan(&position)

	return position, err
}

func (r *Repository) Create(
	ctx context.Context,
	tab *Tab,
) error {

	metadataBytes, err := json.Marshal(
		tab.Metadata,
	)

	if err != nil {
		return err
	}

	query := `
		INSERT INTO room_tabs (
			room_id,
			created_by,
			type,
			title,
			position,
			is_active,
			metadata
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING
			id,
			created_at,
			updated_at
	`

	return r.db.QueryRow(
		ctx,
		query,
		tab.RoomID,
		tab.CreatedBy,
		tab.Type,
		tab.Title,
		tab.Position,
		tab.IsActive,
		metadataBytes,
	).Scan(
		&tab.ID,
		&tab.CreatedAt,
		&tab.UpdatedAt,
	)
}

func (r *Repository) ListByRoomID(
	ctx context.Context,
	roomID string,
) ([]Tab, error) {

	query := `
		SELECT
			id,
			room_id,
			created_by,
			type,
			title,
			position,
			is_active,
			metadata,
			created_at,
			updated_at
		FROM room_tabs
		WHERE room_id = $1
		ORDER BY position ASC
	`

	rows, err := r.db.Query(
		ctx,
		query,
		roomID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tabs []Tab

	for rows.Next() {

		var tab Tab

		var metadataBytes []byte

		err := rows.Scan(
			&tab.ID,
			&tab.RoomID,
			&tab.CreatedBy,
			&tab.Type,
			&tab.Title,
			&tab.Position,
			&tab.IsActive,
			&metadataBytes,
			&tab.CreatedAt,
			&tab.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		if metadataBytes != nil {

			err = json.Unmarshal(
				metadataBytes,
				&tab.Metadata,
			)

			if err != nil {
				return nil, err
			}
		}

		tabs = append(
			tabs,
			tab,
		)
	}

	return tabs, nil
}

func (r *Repository) ClearActiveTabs(
	ctx context.Context,
	roomID string,
) error {

	query := `
		UPDATE room_tabs
		SET is_active = false
		WHERE room_id = $1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		roomID,
	)

	return err
}

func (r *Repository) SetActiveTab(
	ctx context.Context,
	tabID string,
) error {

	query := `
		UPDATE room_tabs
		SET is_active = true
		WHERE id = $1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		tabID,
	)

	return err
}