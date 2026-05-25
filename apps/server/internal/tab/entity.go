package tab

import "time"

type Tab struct {
	ID string `json:"id"`

	RoomID string `json:"room_id"`

	CreatedBy string `json:"created_by"`

	Type string `json:"type"`

	Title string `json:"title"`

	Position int `json:"position"`

	IsActive bool `json:"is_active"`

	Metadata map[string]any `json:"metadata"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}