package app

import (
	"context"
	"gitlab.com/ricardo134/link-service/internal/core/ports"
)

type PartyService interface {
	ports.PartyService
}

type partyService struct {
	notifier ports.PartyNotifier
}

func NewPartyService(n ports.PartyNotifier) PartyService {
	return partyService{n}
}

func (p partyService) Request(ctx context.Context, partyID uint) (any, error) {
	return p.notifier.Requested(ctx, partyID)
}

func (p partyService) Joined(ctx context.Context, partyID uint, userID uint) error {
	return p.notifier.Joined(ctx, partyID, userID)
}
