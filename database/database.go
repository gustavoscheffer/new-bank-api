package database

import (
	"log"
	"new-bank-api/config"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() {
	err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 12 * time.Second}, config.Config("MONGO_DB"), options.Client().ApplyURI(config.Config("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
}
