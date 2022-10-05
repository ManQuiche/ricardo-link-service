package cockroachdb

import (
	"context"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
	"gitlab.com/ricardo134/link-service/internal/core/ports"
	"gorm.io/gorm"
)

type inviteRepository struct {
	client *gorm.DB
}

func NewInviteRepository(client *gorm.DB) ports.InviteRepository {
	return inviteRepository{client: client}
}

func (p inviteRepository) Get(ctx context.Context, inviteID uint) (*entities.Invite, error) {
	var invite entities.Invite
	err := p.client.First(&invite, inviteID).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return &invite, nil
}

func (p inviteRepository) GetAllForUser(ctx context.Context, userID uint) ([]entities.Invite, error) {
	var invites []entities.Invite
	err := p.client.Where("user_id = ?", userID).Find(&invites).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return invites, nil
}

func (p inviteRepository) GetAllForParty(ctx context.Context, partyID uint) ([]entities.Invite, error) {
	var invites []entities.Invite
	err := p.client.Where("party_id = ?", partyID).Find(&invites).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return invites, nil
}

func (p inviteRepository) GetAll(ctx context.Context) ([]entities.Invite, error) {
	var invites []entities.Invite
	err := p.client.Find(&invites).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return invites, nil
}

func (p inviteRepository) Save(ctx context.Context, invite entities.Invite) (*entities.Invite, error) {
	err := p.client.Save(&invite).Error

	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return &invite, err
}

func (p inviteRepository) Delete(ctx context.Context, inviteID uint) error {
	err := p.client.Delete(&entities.Invite{}, inviteID).Error
	if err != nil {
		return notFoundOrElseError(err)
	}

	return nil
}

func (p inviteRepository) DeleteForUser(ctx context.Context, userID uint) error {
	err := p.client.Where("user_id = ?", userID).Delete(&entities.Invite{}).Error
	if err != nil {
		return notFoundOrElseError(err)
	}

	return nil
}

func (p inviteRepository) DeleteForParty(ctx context.Context, partyID uint) error {
	err := p.client.Where("party_id = ?", partyID).Delete(&entities.Invite{}).Error
	if err != nil {
		return notFoundOrElseError(err)
	}

	return nil
}
