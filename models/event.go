// Package models is used to define the structure the models
package models

type Event struct {
	ID               int    `gorm:"primaryKey;autoIncrement" json:"id"`
	EventName        string `gorm:"not null" json:"event_name"`
	Date             string `json:"date"`
	TotalTickets     int    `json:"TotalTickets"`
	RemainingTickets int    `json:"RemainingTickets"`
}

type Booking struct {
	ID            int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerName  string `json:"CustomerName"`
	CustomerPhone string `json:"CustomerPhone"`
	EventID       int    `json:"EventID"`
	BookedTickets int    `json:"BookedTickets"`
}
