package migration

import (
	"log"
	"mpesa/database"
	"mpesa/model"
)

func Migrate() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("Migrating")
	Transaction := &model.Transaction{}
	database.Db.AutoMigrate(&Transaction)
}
