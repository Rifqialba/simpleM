package workspace

import "time"

type Workspace struct {
	ID string `json:"id"`

	OwnerID string `json:"owner_id"`

	Name string `json:"name"`

	Slug string `json:"slug"`

	Description *string `json:"description"`

	IsPersonal bool `json:"is_personal"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}