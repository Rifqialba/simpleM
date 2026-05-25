package user

import "time"

type User struct {
	ID string `json:"id"`

	Email string `json:"email"`

	PasswordHash string `json:"password_hash"`

	Username string `json:"username"`

	DisplayName string `json:"display_name"`

	AvatarURL *string `json:"avatar_url"`

	IsActive bool `json:"is_active"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}