package main

import (
	"banking-app/backend/api"
	"banking-app/backend/configs"
	"banking-app/backend/migrations"
)

func main() {
	configs.InitConfig()
	migrations.Migrate()
	api.StartApi()
}


