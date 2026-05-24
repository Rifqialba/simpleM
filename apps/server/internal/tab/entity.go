package tab

import "time"

type Tab struct {
	ID string

	RoomID string

	CreatedBy string

	Type string

	Title string

	Position int

	IsActive bool

	Metadata map[string]any

	CreatedAt time.Time

	UpdatedAt time.Time
}