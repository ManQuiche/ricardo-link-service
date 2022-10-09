package postgresql

import (
	"context"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
	"gitlab.com/ricardo134/link-service/internal/core/ports"
	"gorm.io/gorm"
)

type inviteRepository struct {
	client *gorm.DB
}

func NewInviteRepository(client *gorm.DB) ports.LinkRepository {
	return inviteRepository{client: client}
}

func (p inviteRepository) Get(ctx context.Context, inviteID uint) (*entities.Link, error) {
	var invite entities.Link
	err := p.client.First(&invite, inviteID).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return &invite, nil
}

func (p inviteRepository) GetAllForParty(ctx context.Context, partyID uint) ([]entities.Link, error) {
	var invites []entities.Link
	err := p.client.Where("party_id = ?", partyID).Find(&invites).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return invites, nil
}

func (p inviteRepository) GetAll(ctx context.Context) ([]entities.Link, error) {
	var invites []entities.Link
	err := p.client.Find(&invites).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return invites, nil
}

func (p inviteRepository) Save(ctx context.Context, invite entities.Link) (*entities.Link, error) {
	err := p.client.Save(&invite).Error

	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return &invite, err
}

func (p inviteRepository) Delete(ctx context.Context, inviteID uint) error {
	err := p.client.Delete(&entities.Link{}, inviteID).Error
	if err != nil {
		return notFoundOrElseError(err)
	}

	return nil
}

func (p inviteRepository) DeleteForParty(ctx context.Context, partyID uint) error {
	err := p.client.Where("party_id = ?", partyID).Delete(&entities.Link{}).Error
	if err != nil {
		return notFoundOrElseError(err)
	}

	return nil
}
