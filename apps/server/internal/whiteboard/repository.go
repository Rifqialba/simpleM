package whiteboard

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

func (r *Repository) Save(
	ctx context.Context,
	whiteboard *Whiteboard,
) error {

	sceneBytes, err := json.Marshal(
		whiteboard.Scene,
	)

	if err != nil {
		return err
	}

	query := `
		INSERT INTO whiteboards (
			tab_id,
			scene,
			version
		)
		VALUES ($1, $2, $3)
		ON CONFLICT (tab_id)
		DO UPDATE SET
			scene = EXCLUDED.scene,
			version = whiteboards.version + 1,
			updated_at = now()
		RETURNING
			id,
			version,
			created_at,
			updated_at
	`

	return r.db.QueryRow(
		ctx,
		query,
		whiteboard.TabID,
		sceneBytes,
		whiteboard.Version,
	).Scan(
		&whiteboard.ID,
		&whiteboard.Version,
		&whiteboard.CreatedAt,
		&whiteboard.UpdatedAt,
	)
}

func (r *Repository) FindByTabID(
	ctx context.Context,
	tabID string,
) (*Whiteboard, error) {

	query := `
		SELECT
			id,
			tab_id,
			scene,
			version,
			created_at,
			updated_at
		FROM whiteboards
		WHERE tab_id = $1
	`

	var whiteboard Whiteboard

	var sceneBytes []byte

	err := r.db.QueryRow(
		ctx,
		query,
		tabID,
	).Scan(
		&whiteboard.ID,
		&whiteboard.TabID,
		&sceneBytes,
		&whiteboard.Version,
		&whiteboard.CreatedAt,
		&whiteboard.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(
		sceneBytes,
		&whiteboard.Scene,
	)

	if err != nil {
		return nil, err
	}

	return &whiteboard, nil
}