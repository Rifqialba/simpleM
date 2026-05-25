package tab

import (
	"context"

	"github.com/gofiber/fiber/v2"
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

func isValidTabType(
	tabType string,
) bool {

	switch tabType {

	case TypeWhiteboard:
		return true

	case TypeMarkdown:
		return true

	case TypePresentation:
		return true

	default:
		return false
	}
}

func (s *Service) Create(
	ctx context.Context,
	roomID string,
	userID string,
	req CreateTabRequest,
) (*Tab, error) {

	if !isValidTabType(req.Type) {

		return nil, fiber.NewError(
			fiber.StatusBadRequest,
			"invalid tab type",
		)
	}

	position, err := s.repo.GetNextPosition(
		ctx,
		roomID,
	)

	if err != nil {
		return nil, err
	}

	tab := &Tab{
		RoomID: roomID,

		CreatedBy: userID,

		Type: req.Type,

		Title: req.Title,

		Position: position,

		IsActive: false,

		Metadata: map[string]any{},
	}

	err = s.repo.Create(
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

	// TODO:
	// wrap active tab updates in transaction
	// to avoid race conditions

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