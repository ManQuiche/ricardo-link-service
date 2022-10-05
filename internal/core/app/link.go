package app

import (
	"context"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
	"gitlab.com/ricardo134/link-service/internal/core/ports"
)

type LinkService interface {
	ports.LinkService
}

type linkService struct {
	repo ports.LinkRepository
}

func NewLinkService(repo ports.LinkRepository) LinkService {
	return linkService{
		repo: repo,
	}
}

func (p linkService) Get(ctx context.Context, linkID uint) (*entities.Link, error) {
	return p.repo.Get(ctx, linkID)
}

func (p linkService) GetAll(ctx context.Context) ([]entities.Link, error) {
	return p.repo.GetAll(ctx)
}

func (p linkService) GetAllForParty(ctx context.Context, partyID uint) ([]entities.Link, error) {
	return p.repo.GetAllForParty(ctx, partyID)
}

func (p linkService) Save(ctx context.Context, link entities.Link) (*entities.Link, error) {
	return p.repo.Save(ctx, link)
}

func (p linkService) Delete(ctx context.Context, linkID uint) error {
	return p.repo.Delete(ctx, linkID)
}

func (p linkService) DeleteForParty(ctx context.Context, partyID uint) error {
	return p.repo.DeleteForParty(ctx, partyID)
}
