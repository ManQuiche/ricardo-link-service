package nats

import (
	"context"
	"gitlab.com/ricardo134/link-service/internal/core/app"
	"gitlab.com/ricardo134/link-service/internal/driving/async"
)

type natsHandler struct {
	linkService app.LinkService
}

func NewNatsLinkHandler(inviteSvc app.LinkService) async.Handler {
	return natsHandler{inviteSvc}
}

func (nh natsHandler) OnUserDelete(userID uint) {
	_ = nh.linkService.DeleteForUser(context.Background(), userID)
}

func (nh natsHandler) OnPartyDelete(partyID uint) {
	_ = nh.linkService.DeleteForParty(context.Background(), partyID)
}
