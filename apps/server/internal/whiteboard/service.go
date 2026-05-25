package whiteboard

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

func (s *Service) Save(
	ctx context.Context,
	tabID string,
	req SaveWhiteboardRequest,
) (*Whiteboard, error) {

	whiteboard := &Whiteboard{
		TabID: tabID,

		Scene: req.Scene,

		Version: 1,
	}

	err := s.repo.Save(
		ctx,
		whiteboard,
	)

	if err != nil {
		return nil, err
	}

	return whiteboard, nil
}

func (s *Service) FindByTabID(
	ctx context.Context,
	tabID string,
) (*Whiteboard, error) {

	return s.repo.FindByTabID(
		ctx,
		tabID,
	)
}