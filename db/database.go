// Package db is used for database connection
package db

import (
	"fmt"

	"github.com/HarshithRajesh/BookIt_Go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}
}

func SeedEvents() {
	var count int64
	// Check if there are any records in the events table
	DB.Model(&models.Event{}).Count(&count)

	if count == 0 {
		events := []models.Event{
			{EventName: "Golang conference", Date: "20 Feb 2026", TotalTickets: 100, RemainingTickets: 100},
			{EventName: "Toxic Movie", Date: "15 March 2026", TotalTickets: 100, RemainingTickets: 100},
			{EventName: "Junior Level Hackathon", Date: "25 April 2026", TotalTickets: 100, RemainingTickets: 100},
			{EventName: "KubCon", Date: "10 May 2026", TotalTickets: 100, RemainingTickets: 100},
			{EventName: "Golang International conference", Date: "20 May 2026", TotalTickets: 100, RemainingTickets: 100},
		}

		if err := DB.Create(&events).Error; err != nil {
			fmt.Printf("Could not seed events: %v\n", err)
			return
		}
	}
}
