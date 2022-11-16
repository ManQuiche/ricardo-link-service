package entities

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
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
	Signature string `json:"magic_link"`
}

func (m MagicLink) String() string {
	jsonL, err := json.Marshal(m.ShortLink)
	if err != nil {
		return ""
	}

	return fmt.Sprintf(
		"%s%s%s",
		base64.URLEncoding.EncodeToString(jsonL),
		magicLinkSep,
		m.Signature,
	)
}

func NewMagicLink(shortL ShortLink, secret []byte) (MagicLink, error) {
	jsonLink, err := json.Marshal(shortL)
	if err != nil {
		return MagicLink{}, errors.New(fmt.Sprintf("could not marshal short link %d: %s", shortL.ID, err))
	}
	digest := sha512.Sum512(append(jsonLink, secret...))

	return MagicLink{
		ShortLink: shortL,
		Signature: base64.URLEncoding.EncodeToString(digest[:]),
	}, nil
}

func NewMagicLinkFromString(magicL string) (MagicLink, error) {
	b64ShortL, sig, found := strings.Cut(magicL, magicLinkSep)
	if !found {
		return MagicLink{}, errors.New("could not decode magic link")
	}

	jsonShortL, err := base64.URLEncoding.DecodeString(b64ShortL)
	if err != nil {
		return MagicLink{}, errors.New("could not decode b64 short link")
	}

	var shortL ShortLink
	err = json.Unmarshal([]byte(jsonShortL), &shortL)
	if err != nil {
		return MagicLink{}, errors.New("could not unmarshal short link")
	}

	return MagicLink{
		ShortLink: shortL,
		Signature: sig,
	}, nil
}

func (m MagicLink) IsValid(secret []byte) (bool, error) {
	jsonLink, err := json.Marshal(m.ShortLink)
	if err != nil {
		return false, errors.New(fmt.Sprintf("could not marshal short link %d: %s", m.ShortLink.ID, err))
	}

	// TODO: just store the previously computed signature already...
	digest := sha512.Sum512(append(jsonLink, secret...))

	return base64.URLEncoding.EncodeToString(digest[:]) == m.Signature, nil
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
