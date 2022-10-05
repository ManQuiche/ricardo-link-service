package entities

import (
	"gorm.io/gorm"
	"time"
)

type Link struct {
	ID         uint           `json:"id" gorm:"primarykey"`
	PartyID    uint           `json:"party_id,omitempty" gorm:"index,notNull"`
	Expiration time.Time      `json:"expiration,omitempty"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type CreateLinkRequest struct {
	PartyID    uint      `json:"party_id,omitempty" binding:"required"`
	Expiration time.Time `json:"expiration,omitempty"`
}

type DeleteInviteRequest struct {
	ID uint
}
