package user

import (
	"context"

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

	user := &User{
		Email:       req.Email,
		Username:    req.Username,
		DisplayName: req.DisplayName,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}