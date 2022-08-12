package orm

import (
	"gin-project-starter/database"
	"log"

	"gorm.io/gorm"
)

var DbClient *gorm.DB

func init() {
	var err error
	DbClient, err = database.Connect()

	if err != nil {
		log.Fatal(err)
	}
}
