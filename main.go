package main

import (
	"github.com/HarshithRajesh/BookIt_Go/cmd"
	"github.com/HarshithRajesh/BookIt_Go/db"
	"github.com/HarshithRajesh/BookIt_Go/models"
)

func main() {
	db.ConnectDB()
	db.DB.AutoMigrate(&models.Event{}, &models.Booking{})
	db.SeedEvents()

	cmd.Execute()
}
