package room

type CreateRoomRequest struct {
	Name string `json:"name" validate:"required,min=3,max=64"`

	Description *string `json:"description"`
}

type RoomResponse struct {
	ID string `json:"id"`

	WorkspaceID string `json:"workspace_id"`

	CreatedBy string `json:"created_by"`

	Name string `json:"name"`

	Description *string `json:"description"`

	IsArchived bool `json:"is_archived"`
}