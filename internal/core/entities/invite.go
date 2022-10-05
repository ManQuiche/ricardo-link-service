package entities

import "gorm.io/gorm"

type Invite struct {
	gorm.Model  `json:"gorm.Model"`
	Explanation string `json:"explanation,omitempty"`
	PartyID     uint   `json:"partyID,omitempty" gorm:"index,notNull"`
	UserID      uint   `json:"userID,omitempty" gorm:"index,notNull"`
	Answered    bool   `json:"answered,omitempty"`
	Accepted    bool   `json:"accepted,omitempty"`
}

type CreateInviteRequest struct {
	PartyID uint `json:"partyID,omitempty" binding:"required"`
	UserID  uint `json:"userID,omitempty" binding:"required"`
}

type UpdateInviteRequest struct {
	Explanation string
	ID          uint
	Answered    bool
	Accepted    bool
}

type DeleteInviteRequest struct {
	ID uint
}
