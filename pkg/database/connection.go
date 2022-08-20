package database

import (
	"gin-project-starter/pkg/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbClient *gorm.DB

func Connect() (*gorm.DB, error) {

	pgConf := postgres.Config{
		DSN:                  env.Get("DATABASE_URL"),
		PreferSimpleProtocol: true,
	}

	db, err := gorm.Open(postgres.New(pgConf), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default,
	})

	if err != nil {
		panic(err)
	}

	return db, err
}
