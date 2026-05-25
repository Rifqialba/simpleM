package room

import "time"

type Room struct {
	ID string `json:"id"`

	WorkspaceID string `json:"workspace_id"`

	CreatedBy string `json:"created_by"`

	Name string `json:"name"`

	Description *string `json:"description"`

	IsArchived bool `json:"is_archived"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}