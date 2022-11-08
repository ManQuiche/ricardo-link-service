package link

import (
	"context"
	"errors"
	"fmt"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
)

func (p service) Get(ctx context.Context, linkID uint) (*entities.Link, error) {
	link, err := p.repo.Get(ctx, linkID)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (p service) GetMagic(ctx context.Context, linkID uint) (*entities.MagicLink, error) {
	link, err := p.Get(ctx, linkID)
	if err != nil {
		return nil, err
	}
	mLink, err := p.ToMagic(ctx, *link)
	if err != nil {
		return nil, err
	}
	return mLink, nil
}

func (p service) GetAll(ctx context.Context) ([]entities.Link, error) {
	return p.repo.GetAll(ctx)
}

func (p service) GetAllForParty(ctx context.Context, partyID uint) ([]entities.Link, error) {
	return p.repo.GetAllForParty(ctx, partyID)
}

func (p service) Save(ctx context.Context, link entities.Link) (*entities.Link, error) {
	l, err := p.repo.Save(ctx, link)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not save link %d: %s", link.ID, err))
	}

	magicLink, err := p.ToMagic(ctx, link)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not convert to magic %d: %s", link.ID, err))
	}

	_, err = p.extlink.Create(ctx, magicLink.String(), l.ID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not create ext link %d: %s", link.ID, err))
	}

	return l, nil
}

func (p service) Delete(ctx context.Context, linkID uint) error {
	return p.repo.Delete(ctx, linkID)
}

func (p service) DeleteForParty(ctx context.Context, partyID uint) error {
	return p.repo.DeleteForParty(ctx, partyID)
}

func (p service) DeleteForUser(ctx context.Context, userID uint) error {
	return p.repo.DeleteForUser(ctx, userID)
}
