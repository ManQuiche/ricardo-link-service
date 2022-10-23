package ports

import "context"

type PartyService interface {
	Request(ctx context.Context, partyID uint) (any, error)
	Joined(ctx context.Context, partyID uint, userID uint) error
}

// PartyNotifier Interface designed as a template for event
// publishing function on party events
type PartyNotifier interface {
	Requested(ctx context.Context, partyID uint) (any, error)
	Joined(ctx context.Context, partyID uint, userID uint) error
}
