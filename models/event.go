package models

type Event struct {
	ID               int
	EventName        string
	Date             string
	TotalTickets     int
	RemainingTickets int
}

type Booking struct {
	ID            int
	CustomerName  string
	CustomerPhone string
	EventID       int
	BookedTickets int
}
