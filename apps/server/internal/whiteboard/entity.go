package whiteboard

import "time"

type Whiteboard struct {
	ID string `json:"id"`

	TabID string `json:"tab_id"`

	Scene map[string]any `json:"scene"`

	Version int64 `json:"version"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}