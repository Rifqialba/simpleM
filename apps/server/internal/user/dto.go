package user

type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email"`

	Username string `json:"username" validate:"required,min=3,max=32"`

	Password string `json:"password" validate:"required,min=6"`

	DisplayName string `json:"display_name" validate:"required,min=3,max=64"`
}

type UserResponse struct {
	ID string `json:"id"`

	Email string `json:"email"`

	Username string `json:"username"`
	

	DisplayName string `json:"display_name"`

	AvatarURL *string `json:"avatar_url"`

	IsActive bool `json:"is_active"`
}

type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`

	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`

	User UserResponse `json:"user"`
}