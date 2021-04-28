package main

import (
	"go-bank-backend/api"
	"go-bank-backend/migrations"
)

func main() {
	migrations.Migrate()
	api.StartApi()
}
