package storage

import (
	"github.com/HarshithRajesh/BookIt_Go/models"
)

var events = []models.Event{
	{ID: 1, EventName: "Golang conference", Date: "20 Feb 2026", TotalTickets: 100, RemainingTickets: 100},
	{ID: 2, EventName: "Toxic Movie", Date: "15 March 2026", TotalTickets: 100, RemainingTickets: 100},
	{ID: 3, EventName: "Junior Level Hackathon", Date: "25 April 2026", TotalTickets: 100, RemainingTickets: 100},
	{ID: 4, EventName: "KubCon", Date: "10 May 2026", TotalTickets: 100, RemainingTickets: 100},
	{ID: 5, EventName: "Golang International conference", Date: "20 May 2026", TotalTickets: 100, RemainingTickets: 100}}

var bookings = []models.Booking{}
