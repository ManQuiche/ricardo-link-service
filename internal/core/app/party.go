package app

import (
	"context"
	"gitlab.com/ricardo134/link-service/internal/core/ports"
)

type PartyService interface {
	ports.PartyService
}

type partyService struct {
	requestor ports.PartyRequestor
}

func NewPartyService(r ports.PartyRequestor) PartyService {
	return partyService{r}
}

func (p partyService) Request(ctx context.Context, partyID uint) (any, error) {
	return p.requestor.Requested(ctx, partyID)
}
