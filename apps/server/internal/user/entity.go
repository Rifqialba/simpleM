package user

import "time"

type User struct {
	ID string

	Email string

	PasswordHash string

	Username string

	DisplayName string

	AvatarURL *string

	IsActive bool

	CreatedAt time.Time

	UpdatedAt time.Time
}