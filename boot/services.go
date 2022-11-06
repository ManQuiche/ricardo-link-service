package boot

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"gitlab.com/ricardo134/link-service/internal/core/app/link"
	"gitlab.com/ricardo134/link-service/internal/core/app/party"
	"log"

	natsextout "gitlab.com/ricardo134/link-service/internal/driven/async/nats"
	"gitlab.com/ricardo134/link-service/internal/driven/db/postgresql"
	"gitlab.com/ricardo134/link-service/internal/driving/async"
	natsextin "gitlab.com/ricardo134/link-service/internal/driving/async/nats"
)

var (
	linkService  link.Service
	partyService party.PartyService

	natsEncConn  *nats.EncodedConn
	asyncHandler async.Handler
)

func LoadServices() {
	natsConn, err := nats.Connect(fmt.Sprintf("nats://%s:%s@%s", natsUsr, natsPwd, natsURL))
	if err != nil {
		log.Fatal(err)
	}
	natsEncConn, err = nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)

	linkRepo := postgresql.NewInviteRepository(client)
	linkService = link.NewService(linkRepo, []byte(linkSecret))

	partyNotifier := natsextout.NewPartyNotifier(natsEncConn, natsPartyRequested, natsPartyJoined)
	partyService = party.NewPartyService(partyNotifier)

	asyncHandler = natsextin.NewNatsLinkHandler(linkService)
}
