package workspace

import "time"

type Workspace struct {
	ID string

	OwnerID string

	Name string

	Slug string

	Description *string

	IsPersonal bool

	CreatedAt time.Time

	UpdatedAt time.Time
}