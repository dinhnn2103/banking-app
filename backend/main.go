package main

import (
	"banking-app/backend/api"
	"banking-app/backend/configs"
	"banking-app/backend/database"
)

func main() {
	configs.InitConfig()
	database.InitDB()
	api.StartApi()
}


