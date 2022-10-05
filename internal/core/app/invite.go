package app

import (
	"context"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
	"gitlab.com/ricardo134/link-service/internal/core/ports"
)

type InviteService interface {
	ports.InviteService
}

type inviteService struct {
	repo ports.InviteRepository
}

func NewInviteService(repo ports.InviteRepository) InviteService {
	return inviteService{
		repo: repo,
	}
}

func (p inviteService) Get(ctx context.Context, inviteID uint) (*entities.Invite, error) {
	return p.repo.Get(ctx, inviteID)
}

func (p inviteService) GetAll(ctx context.Context) ([]entities.Invite, error) {
	return p.repo.GetAll(ctx)
}

func (p inviteService) GetAllForUser(ctx context.Context, userID uint) ([]entities.Invite, error) {
	return p.repo.GetAllForUser(ctx, userID)
}

func (p inviteService) GetAllForParty(ctx context.Context, partyID uint) ([]entities.Invite, error) {
	return p.repo.GetAllForParty(ctx, partyID)
}

func (p inviteService) Save(ctx context.Context, invite entities.Invite) (*entities.Invite, error) {
	return p.repo.Save(ctx, invite)
}

func (p inviteService) Delete(ctx context.Context, inviteID uint) error {
	return p.repo.Delete(ctx, inviteID)
}

func (p inviteService) DeleteForUser(ctx context.Context, userID uint) error {
	return p.repo.DeleteForUser(ctx, userID)
}

func (p inviteService) DeleteForParty(ctx context.Context, partyID uint) error {
	return p.repo.DeleteForParty(ctx, partyID)
}
