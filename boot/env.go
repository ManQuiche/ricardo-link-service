package boot

import (
	"gitlab.com/ricardo134/link-service/pkg/errors"
	"log"
	"os"
	"strconv"
)

var (
	dbHost       string
	dbPort       string
	dbUser       string
	dbPassword   string
	dbDatabase   string
	port         string
	url          string
	accessSecret string

	natsURL            string
	natsUsr            string
	natsPwd            string
	natsPartyDeleted   string
	natsPartyRequested string
	natsPartyJoined    string

	linkSecret string

	firebaseKeyFile string
)

func LoadEnv() {
	dbHost = env("DB_HOST")
	dbPort = env("DB_PORT")
	dbUser = env("DB_USER")
	dbDatabase = env("DB_DATABASE")
	dbPassword = env("DB_PASSWORD")

	port = env("PORT")
	url = env("URL")
	accessSecret = env("ACCESS_SECRET")

	natsURL = env("NATS_URL")
	natsUsr = env("NATS_USR")
	natsPwd = env("NATS_PWD")
	natsPartyDeleted = env("NATS_PARTY_DELETED")
	natsPartyRequested = env("NATS_PARTY_REQUESTED")
	natsPartyJoined = env("NATS_USER_JOINED")

	linkSecret = env("LINK_SECRET")

	firebaseKeyFile = env("FIREBASE_KEY_FILE")
}

func envBool(name string) bool {
	res, err := strconv.ParseBool(env(name))
	if err != nil {
		log.Fatalf("env var %s needs to be of boolean type", name)
	}

	return res
}

func env(name string) string {
	str, ok := os.LookupEnv(name)
	if !ok {
		errors.MissingEnvVarF(name)
	}

	return str
}
