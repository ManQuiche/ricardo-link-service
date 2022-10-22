package ports

import (
	"context"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
)

type LinkService interface {
	Get(ctx context.Context, inviteID uint) (*entities.MagicLink, error)
	GetAllForParty(ctx context.Context, partyID uint) ([]entities.Link, error)
	GetAll(ctx context.Context) ([]entities.Link, error)
	Save(ctx context.Context, Link entities.Link) (*entities.MagicLink, error)
	Delete(ctx context.Context, LinkID uint) error
	DeleteForParty(ctx context.Context, partyID uint) error
	DeleteForUser(ctx context.Context, userID uint) error
	IsValid(ctx context.Context, m entities.MagicLink) (bool, error)
}

//type LinkRepository interface {
//	Get(ctx context.Context, LinkID uint) (*entities.Link, error)
//	GetAllForUser(ctx context.Context, userID uint) ([]entities.Link, error)
//	GetAllForParty(ctx context.Context, partyID uint) ([]entities.Link, error)
//	GetAll(ctx context.Context) ([]entities.Link, error)
//	Save(ctx context.Context, link entities.Link) (*entities.Link, error)
//	Delete(ctx context.Context, LinkID uint) error
//}

type LinkRepository interface {
	Get(ctx context.Context, inviteID uint) (*entities.Link, error)
	GetAllForParty(ctx context.Context, partyID uint) ([]entities.Link, error)
	GetAll(ctx context.Context) ([]entities.Link, error)
	Save(ctx context.Context, Link entities.Link) (*entities.Link, error)
	Delete(ctx context.Context, LinkID uint) error
	DeleteForParty(ctx context.Context, partyID uint) error
	DeleteForUser(ctx context.Context, userID uint) error
}
