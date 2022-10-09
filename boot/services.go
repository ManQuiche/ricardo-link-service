package boot

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"gitlab.com/ricardo134/link-service/internal/core/app"
	"log"

	"gitlab.com/ricardo134/link-service/internal/driven/db/postgresql"
	"gitlab.com/ricardo134/link-service/internal/driving/async"
	ricardoNats "gitlab.com/ricardo134/link-service/internal/driving/async/nats"
)

var (
	linkService app.LinkService

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
	linkService = app.NewLinkService(linkRepo)

	asyncHandler = ricardoNats.NewNatsInviteHandler(linkService)
}
