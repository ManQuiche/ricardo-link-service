package entities

import (
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMagicLinkCreation(t *testing.T) {
	shortL := ShortLink{
		ID:        0,
		PartyID:   5,
		CreatorID: 2,
	}

	sec := make([]byte, 20)
	rand.Read(sec)

	magicL, err := NewMagicLink(shortL, sec)
	assert.Nil(t, err)
	assert.NotNil(t, magicL)
}

func TestMagicLinkIsValid(t *testing.T) {
	shortL := ShortLink{
		ID:        0,
		PartyID:   5,
		CreatorID: 2,
	}

	sec := make([]byte, 20)
	_, err := rand.Read(sec)
	assert.Nil(t, err)

	magicL, err := NewMagicLink(shortL, sec)
	assert.Nil(t, err)

	valid, err := magicL.IsValid(sec)
	assert.Nil(t, err)
	assert.True(t, valid)
}

func TestMagicLinkIsNotValid(t *testing.T) {
	shortL := ShortLink{
		ID:        0,
		PartyID:   5,
		CreatorID: 2,
	}

	sec := make([]byte, 20)
	_, err := rand.Read(sec)
	assert.Nil(t, err)

	magicL, err := NewMagicLink(shortL, sec)
	assert.Nil(t, err)

	sec2 := make([]byte, 20)
	_, err = rand.Read(sec2)
	assert.Nil(t, err)

	valid, err := magicL.IsValid(sec2)
	assert.Nil(t, err)
	assert.False(t, valid)
}
