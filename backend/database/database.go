package database

import (
	"banking-app/backend/configs"
	"banking-app/backend/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBInfo string
	DBconnection *gorm.DB
	dbConnectError error
)

func InitDB() {
	DBconnection, dbConnectError = gorm.Open(postgres.Open(configs.DBInfo), &gorm.Config{})
	helpers.HandleErr(dbConnectError)
}

