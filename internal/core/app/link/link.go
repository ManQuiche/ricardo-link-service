package link

import (
	"context"
	"errors"
	"fmt"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
	"gitlab.com/ricardo134/link-service/internal/core/ports"
)

type Service interface {
	ports.LinkService
}

type service struct {
	repo       ports.LinkRepository
	extlink    ports.ExternalLinkService
	extlinkURL string
	secret     []byte
}

func NewService(repo ports.LinkRepository, extlink ports.ExternalLinkService, extlinkURL string, secret []byte) Service {
	return service{
		repo:       repo,
		extlink:    extlink,
		extlinkURL: extlinkURL,
		secret:     secret,
	}
}

func (p service) IsValid(ctx context.Context, m entities.MagicLink) (bool, error) {
	return m.IsValid(p.secret)
}

func (p service) ToMagic(ctx context.Context, link entities.Link) (*entities.MagicLink, error) {
	shortL := entities.ShortLink{
		ID:        link.ID,
		PartyID:   link.PartyID,
		CreatorID: link.CreatorID,
	}

	magicLink, err := entities.NewMagicLink(shortL, p.secret)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("toMagic: %s", err))
	}

	return &magicLink, nil
}
