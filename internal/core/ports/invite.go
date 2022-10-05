package ports

import (
	"context"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
)

type InviteService interface {
	Get(ctx context.Context, inviteID uint) (*entities.Invite, error)
	GetAllForUser(ctx context.Context, userID uint) ([]entities.Invite, error)
	GetAllForParty(ctx context.Context, partyID uint) ([]entities.Invite, error)
	GetAll(ctx context.Context) ([]entities.Invite, error)
	Save(ctx context.Context, invite entities.Invite) (*entities.Invite, error)
	Delete(ctx context.Context, inviteID uint) error
	DeleteForUser(ctx context.Context, userID uint) error
	DeleteForParty(ctx context.Context, partyID uint) error
}

//type InviteRepository interface {
//	Get(ctx context.Context, inviteID uint) (*entities.Invite, error)
//	GetAllForUser(ctx context.Context, userID uint) ([]entities.Invite, error)
//	GetAllForParty(ctx context.Context, partyID uint) ([]entities.Invite, error)
//	GetAll(ctx context.Context) ([]entities.Invite, error)
//	Save(ctx context.Context, invite entities.Invite) (*entities.Invite, error)
//	Delete(ctx context.Context, inviteID uint) error
//}

type InviteRepository interface {
	InviteService
}
