package entities

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MagicLink_Creation(t *testing.T) {
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

func Test_MagicLink_IsValid(t *testing.T) {
	shortL := ShortLink{
		ID:        80,
		PartyID:   1,
		CreatorID: 9000,
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

func Test_MagicLink_IsNotValid(t *testing.T) {
	shortL := ShortLink{
		ID:        35,
		PartyID:   108,
		CreatorID: 40,
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

func Test_MagicLink_String(t *testing.T) {
	shortL := ShortLink{
		ID:        10,
		PartyID:   6,
		CreatorID: 8,
	}

	sec := make([]byte, 20)
	_, err := rand.Read(sec)
	assert.Nil(t, err)

	jsonL, err := json.Marshal(shortL)
	assert.Nil(t, err)

	m, err := NewMagicLink(shortL, sec)
	assert.Nil(t, err)

	want := fmt.Sprintf(
		"%s%s%s",
		base64.URLEncoding.EncodeToString(jsonL),
		magicLinkSep,
		m.Signature,
	)

	assert.Equal(t, want, m.String())
}

func Test_MagicLink_FromString(t *testing.T) {
	shortL := ShortLink{
		ID:        2,
		PartyID:   3,
		CreatorID: 5,
	}

	sec := make([]byte, 20)
	_, err := rand.Read(sec)
	assert.Nil(t, err)

	jsonL, err := json.Marshal(shortL)
	assert.Nil(t, err)

	wantM, err := NewMagicLink(shortL, sec)
	assert.Nil(t, err)

	properLink := fmt.Sprintf(
		"%s%s%s",
		base64.URLEncoding.EncodeToString(jsonL),
		magicLinkSep,
		wantM.Signature,
	)

	fromLink, err := NewMagicLinkFromString(properLink)
	assert.Nil(t, err)

	assert.Equal(t, wantM, fromLink)
}
