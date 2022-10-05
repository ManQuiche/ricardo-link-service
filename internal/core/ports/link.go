package ports

import (
	"context"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
)

type LinkService interface {
	Get(ctx context.Context, inviteID uint) (*entities.Link, error)
	GetAllForParty(ctx context.Context, partyID uint) ([]entities.Link, error)
	GetAll(ctx context.Context) ([]entities.Link, error)
	Save(ctx context.Context, invite entities.Link) (*entities.Link, error)
	Delete(ctx context.Context, inviteID uint) error
	DeleteForParty(ctx context.Context, partyID uint) error
}

//type LinkRepository interface {
//	Get(ctx context.Context, inviteID uint) (*entities.Link, error)
//	GetAllForUser(ctx context.Context, userID uint) ([]entities.Link, error)
//	GetAllForParty(ctx context.Context, partyID uint) ([]entities.Link, error)
//	GetAll(ctx context.Context) ([]entities.Link, error)
//	Save(ctx context.Context, invite entities.Link) (*entities.Link, error)
//	Delete(ctx context.Context, inviteID uint) error
//}

type LinkRepository interface {
	LinkService
}
