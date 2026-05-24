package user

import "time"

type User struct {
	ID string

	Email string

	Username string

	DisplayName string

	AvatarURL *string

	IsActive bool

	CreatedAt time.Time

	UpdatedAt time.Time
}