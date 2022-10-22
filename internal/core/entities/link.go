package entities

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

const (
	magicLinkSep = "."
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

type ShortLink struct {
	ID        uint `json:"id"`
	PartyID   uint `json:"party_id"`
	CreatorID uint `json:"creator_id"`
}

type MagicLink struct {
	ShortLink
	MagicLink string `json:"magic_link"`
}

func (m MagicLink) String() string {
	jsonL, err := json.Marshal(m.ShortLink)
	if err != nil {
		return ""
	}

	return string(jsonL) + magicLinkSep + m.MagicLink
}

func NewMagicLinkFromString(mLink string) (MagicLink, error) {
	jsonShortL, digest, found := strings.Cut(mLink, magicLinkSep)
	if !found {
		return MagicLink{}, errors.New("could not decode magic link")
	}

	var shortL ShortLink
	err := json.Unmarshal([]byte(jsonShortL), &shortL)
	if err != nil {
		return MagicLink{}, errors.New("could not unmarshal short link")
	}

	return MagicLink{
		ShortLink: shortL,
		MagicLink: digest,
	}, nil
}

//func (m MagicLink) MarshalJSON() ([]byte, error) {
//	jsonL, err := json.Marshal(m.ShortLink)
//	if err != nil {
//		return nil, err
//	}
//
//	return append(jsonL, []byte(m.MagicLink)...), nil
//}
//
//func (m MagicLink) UnmarshalJSON(bytes []byte) error {
//
//}

type CreateLinkRequest struct {
	PartyID    uint       `json:"party_id" binding:"required"`
	Expiration *time.Time `json:"expiration,omitempty"`
}

type UpdateLinkRequest struct {
	Expiration *time.Time `json:"expiration"`
}
