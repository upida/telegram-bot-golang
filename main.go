package main

import (
	"bot/models"
	"bot/routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run()
}
