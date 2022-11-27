package boot

import (
	"context"
	"google.golang.org/api/option"
	"log"

	"google.golang.org/api/firebasedynamiclinks/v1"
)

var (
	fbLinkService *firebasedynamiclinks.Service
)

func InitFirebaseApp() {
	var err error
	opt := option.WithCredentialsFile("boot/ricardo-9b5d5-firebase-adminsdk-udnxf-2e4b3b051f.json")

	fbLinkService, err = firebasedynamiclinks.NewService(context.Background(), opt)
	if err != nil {
		log.Fatalf("firebase dynamic link init: %s", err)
	}
}
