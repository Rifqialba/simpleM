package tab

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(
	repo *Repository,
) *Service {

	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(
	ctx context.Context,
	roomID string,
	userID string,
	req CreateTabRequest,
) (*Tab, error) {

	tab := &Tab{
		RoomID: roomID,

		CreatedBy: userID,

		Type: req.Type,

		Title: req.Title,

		Position: 0,

		IsActive: false,

		Metadata: map[string]any{},
	}

	err := s.repo.Create(
		ctx,
		tab,
	)

	if err != nil {
		return nil, err
	}

	return tab, nil
}

func (s *Service) ListByRoomID(
	ctx context.Context,
	roomID string,
) ([]Tab, error) {

	return s.repo.ListByRoomID(
		ctx,
		roomID,
	)
}

func (s *Service) ActivateTab(
	ctx context.Context,
	roomID string,
	tabID string,
) error {

	err := s.repo.ClearActiveTabs(
		ctx,
		roomID,
	)

	if err != nil {
		return err
	}

	return s.repo.SetActiveTab(
		ctx,
		tabID,
	)
}