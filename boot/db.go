package boot

import (
	"fmt"
	_ "github.com/lib/pq"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
	"gitlab.com/ricardo134/link-service/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var (
	client *gorm.DB
)

func LoadDb() {
	// TODO: finish that shiiit
	//connectionString := fmt.Sprintf("")

	var err error
	client, err = gorm.Open(postgres.Open(
		fmt.Sprint("postgres://", dbUser, ":", dbPassword, "@", dbHost, ":", dbPort, "/", dbDatabase, "?sslmode=disable")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		errors.CannotConnectToDb(dbHost, dbPort)
	}

	err = client.AutoMigrate(&entities.ExternalLink{})
	if err != nil {
		log.Fatal("could not migrate db, exiting...")
	}
}
