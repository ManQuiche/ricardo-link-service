package nats

import (
	"context"
	"gitlab.com/ricardo134/link-service/internal/core/app"
	"gitlab.com/ricardo134/link-service/internal/driving/async"
)

type natsHandler struct {
	inviteService app.InviteService
}

func NewNatsInviteHandler(inviteSvc app.InviteService) async.Handler {
	return natsHandler{inviteSvc}
}

func (nh natsHandler) OnUserDelete(userID uint) {
	_ = nh.inviteService.DeleteForUser(context.Background(), userID)
}

func (nh natsHandler) OnPartyDelete(partyID uint) {
	_ = nh.inviteService.DeleteForParty(context.Background(), partyID)
}
