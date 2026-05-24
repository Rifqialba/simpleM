package workspace

type CreateWorkspaceRequest struct {
	Name string `json:"name" validate:"required,min=3,max=64"`

	Slug string `json:"slug" validate:"required,min=3,max=64"`
}

type WorkspaceResponse struct {
	ID string `json:"id"`

	OwnerID string `json:"owner_id"`

	Name string `json:"name"`

	Slug string `json:"slug"`

	Description *string `json:"description"`

	IsPersonal bool `json:"is_personal"`
}