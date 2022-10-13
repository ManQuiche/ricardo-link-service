package entities

import (
	"gorm.io/gorm"
	"time"
)

type Link struct {
	ID         uint           `json:"id" gorm:"primarykey"`
	PartyID    uint           `json:"party_id" gorm:"index,notNull"`
	CreatorID  uint           `json:"creator_id" gorm:"notNull"`
	Expiration *time.Time     `json:"expiration,omitempty"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type MagicLink struct {
	Link
	MagicLink string `json:"magic_link"`
}

type CreateLinkRequest struct {
	PartyID    uint       `json:"party_id" binding:"required"`
	Expiration *time.Time `json:"expiration,omitempty"`
}

type UpdateLinkRequest struct {
	Expiration *time.Time `json:"expiration"`
}
