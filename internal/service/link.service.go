package service

import (
	"context"

	"github.com/fatawaimamalmuftin/be-shortlink/internal/model"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/repo"
)

type LinkService struct {
	Repo *repo.LinkRepo
}

func NewLinkService(r *repo.LinkRepo) *LinkService {
	return &LinkService{Repo: r}
}

func (s *LinkService) GetMyLinks(ctx context.Context, userID int) ([]model.ResLink, error) {
	links, err := s.Repo.GetLinksByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return links, nil
}
