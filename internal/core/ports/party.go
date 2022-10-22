package ports

import "context"

type PartyService interface {
	Request(ctx context.Context, partyID uint) (any, error)
}

// PartyRequestor Interface designed as a template for event
// publishing function on user registration
type PartyRequestor interface {
	Requested(ctx context.Context, partyID uint) (any, error)
}
