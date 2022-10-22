package nats

import (
	"context"
	"github.com/nats-io/nats.go"
	"gitlab.com/ricardo134/link-service/internal/core/ports"
)

type partyRequestor struct {
	conn           *nats.EncodedConn
	requestedTopic string
}

func NewPartyRequestor(conn *nats.EncodedConn, requestedTopic string) ports.PartyRequestor {
	return partyRequestor{conn, requestedTopic}
}

func (p partyRequestor) Requested(ctx context.Context, partyID uint) (any, error) {
	var party any
	err := p.conn.Request(p.requestedTopic, partyID, &party, nats.DefaultTimeout*2)
	if err != nil {
		return nil, err
	}

	return party, nil
}
