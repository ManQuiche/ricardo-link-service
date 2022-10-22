package app

import (
	"bytes"
	"context"
	"crypto/sha512"
	"encoding/json"
	"github.com/pkg/errors"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
	"gitlab.com/ricardo134/link-service/internal/core/ports"
)

type LinkService interface {
	ports.LinkService
}

type linkService struct {
	repo   ports.LinkRepository
	secret []byte
}

func NewLinkService(repo ports.LinkRepository, secret []byte) LinkService {
	return linkService{
		repo:   repo,
		secret: secret,
	}
}

func (p linkService) Get(ctx context.Context, linkID uint) (*entities.MagicLink, error) {
	link, err := p.repo.Get(ctx, linkID)
	if err != nil {
		return nil, err
	}
	mLink, err := p.toMagic(ctx, *link)
	if err != nil {
		return nil, err
	}
	return mLink, nil
}

func (p linkService) GetAll(ctx context.Context) ([]entities.Link, error) {
	return p.repo.GetAll(ctx)
}

func (p linkService) GetAllForParty(ctx context.Context, partyID uint) ([]entities.Link, error) {
	return p.repo.GetAllForParty(ctx, partyID)
}

func (p linkService) Save(ctx context.Context, link entities.Link) (*entities.MagicLink, error) {
	l, err := p.repo.Save(ctx, link)
	if err != nil {
		return nil, errors.Wrapf(err, "could not save link %d", link.ID)
	}

	return p.toMagic(ctx, *l)
}

func (p linkService) Delete(ctx context.Context, linkID uint) error {
	return p.repo.Delete(ctx, linkID)
}

func (p linkService) DeleteForParty(ctx context.Context, partyID uint) error {
	return p.repo.DeleteForParty(ctx, partyID)
}

func (p linkService) DeleteForUser(ctx context.Context, userID uint) error {
	return p.repo.DeleteForUser(ctx, userID)
}

func (p linkService) IsValid(ctx context.Context, m entities.MagicLink) (bool, error) {
	jsonLink, err := json.Marshal(m.ShortLink)
	if err != nil {
		return false, errors.Wrapf(err, "could not marshal short link %d", m.ShortLink.ID)
	}
	digest := sha512.Sum512(append(jsonLink, p.secret...))

	return bytes.Equal(digest[:], []byte(m.MagicLink)), nil
}

func (p linkService) toMagic(ctx context.Context, link entities.Link) (*entities.MagicLink, error) {
	shortL := entities.ShortLink{
		ID:        link.ID,
		PartyID:   link.PartyID,
		CreatorID: link.CreatorID,
	}

	jsonLink, err := json.Marshal(shortL)
	if err != nil {
		return nil, errors.Wrapf(err, "could not marshal short link %d", link.ID)
	}
	digest := sha512.Sum512(append(jsonLink, p.secret...))

	return &entities.MagicLink{
		ShortLink: shortL,
		MagicLink: string(digest[:]),
	}, nil
}
