package nats

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
	"gitlab.com/ricardo134/link-service/internal/core/ports"
)

type partyNotifier struct {
	conn           *nats.EncodedConn
	requestedTopic string
	joinedTopic    string
}

func NewPartyNotifier(conn *nats.EncodedConn, requestedTopic, joinedTopic string) ports.PartyNotifier {
	return partyNotifier{conn, requestedTopic, joinedTopic}
}

func (p partyNotifier) Requested(ctx context.Context, partyID uint) (any, error) {
	var party any
	err := p.conn.Request(p.requestedTopic, partyID, &party, nats.DefaultTimeout*2)
	if err != nil {
		return nil, err
	}

	return party, nil
}

func (p partyNotifier) Joined(ctx context.Context, partyID uint, userID uint) error {
	fmt.Println("halloa")

	return p.conn.Publish(p.joinedTopic, entities.JoinInfo{PartyID: partyID, UserID: userID})
}
