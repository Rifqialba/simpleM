package room

import "context"

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
	workspaceID string,
	userID string,
	req CreateRoomRequest,
) (*Room, error) {

	room := &Room{
		WorkspaceID: workspaceID,
		CreatedBy: userID,
		Name: req.Name,
		Description: req.Description,
	}

	err := s.repo.Create(
		ctx,
		room,
	)

	if err != nil {
		return nil, err
	}

	return room, nil
}