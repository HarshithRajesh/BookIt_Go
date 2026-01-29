package main

import (
	"github.com/HarshithRajesh/BookIt_Go/cmd"
	"github.com/HarshithRajesh/BookIt_Go/db"
	"github.com/HarshithRajesh/BookIt_Go/models"
)

func main() {
	db.ConnectDB()
	err := db.DB.AutoMigrate(&models.Event{}, &models.Booking{})
	if err != nil {
		panic("couldnt migrate")
	}
	db.SeedEvents()

	cmd.Execute()
}
