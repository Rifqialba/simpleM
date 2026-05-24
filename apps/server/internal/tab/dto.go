package tab

type CreateTabRequest struct {
	Type string `json:"type" validate:"required"`

	Title string `json:"title" validate:"required,min=1,max=128"`
}

type TabResponse struct {
	ID string `json:"id"`

	RoomID string `json:"room_id"`

	CreatedBy string `json:"created_by"`

	Type string `json:"type"`

	Title string `json:"title"`

	Position int `json:"position"`

	IsActive bool `json:"is_active"`
}