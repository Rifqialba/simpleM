package workspace

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
	userID string,
	req CreateWorkspaceRequest,
) (*Workspace, error) {

	workspace := &Workspace{
		OwnerID: userID,
		Name: req.Name,
		Slug: req.Slug,
		IsPersonal: false,
	}

	err := s.repo.Create(
		ctx,
		workspace,
	)

	if err != nil {
		return nil, err
	}

	return workspace, nil
}