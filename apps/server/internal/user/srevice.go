package user

import (
	"context"

	"github.com/Rifqialba/simplem/apps/server/internal/auth"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	repo      *Repository
	validator *validator.Validate
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo:      repo,
		validator: validator.New(),
	}
}

func (s *Service) Create(
	ctx context.Context,
	req CreateUserRequest,
) (*User, error) {

	if err := s.validator.Struct(req); err != nil {
		return nil, err
	}

	hashedPassword, err := auth.HashPassword(
		req.Password,
	)

	if err != nil {
		return nil, err
	}

	user := &User{
		Email:        req.Email,
		Username:     req.Username,
		DisplayName:  req.DisplayName,
		PasswordHash: hashedPassword,
	}

	if err := s.repo.Create(
		ctx,
		user,
	); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Login(
	ctx context.Context,
	req LoginRequest,
	jwtSecret string,
) (*AuthResponse, error) {

	if err := s.validator.Struct(req); err != nil {
		return nil, err
	}

	user, err := s.repo.FindByEmail(
		ctx,
		req.Email,
	)

	if err != nil {
		return nil, err
	}

	if err := auth.CheckPassword(
		req.Password,
		user.PasswordHash,
	); err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(
		user.ID,
		jwtSecret,
	)

	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User: UserResponse{
			ID:          user.ID,
			Email:       user.Email,
			Username:    user.Username,
			DisplayName: user.DisplayName,
			AvatarURL:   user.AvatarURL,
			IsActive:    user.IsActive,
		},
	}, nil
}