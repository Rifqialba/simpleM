package room

import "time"

type Room struct {
	ID string

	WorkspaceID string

	CreatedBy string

	Name string

	Description *string

	IsArchived bool

	CreatedAt time.Time

	UpdatedAt time.Time
}