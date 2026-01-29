// Package storage path implements bussiness logics
package storage

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/HarshithRajesh/BookIt_Go/models"
)

var events = []models.Event{
	{ID: 1, EventName: "Golang conference", Date: "20 Feb 2026", TotalTickets: 100, RemainingTickets: 100},
	{ID: 2, EventName: "Toxic Movie", Date: "15 March 2026", TotalTickets: 100, RemainingTickets: 100},
	{ID: 3, EventName: "Junior Level Hackathon", Date: "25 April 2026", TotalTickets: 100, RemainingTickets: 100},
	{ID: 4, EventName: "KubCon", Date: "10 May 2026", TotalTickets: 100, RemainingTickets: 100},
	{ID: 5, EventName: "Golang International conference", Date: "20 May 2026", TotalTickets: 100, RemainingTickets: 100},
}

var booking = []models.Booking{}

func GetEvents() []string {
	var names []string
	for _, event := range events {
		names = append(names, event.EventName)
	}
	return names
}

func CheckEventTickets(eventID1 string, BookedTickets1 string) error {
	eventID, err := strconv.Atoi(eventID1)
	if err != nil {
		log.Fatal("Couldnt convert Event id to int")
	}
	BookedTickets, err := strconv.Atoi(BookedTickets1)
	if err != nil {
		log.Fatal("Couldnt convert BookedTickets to int")
	}
	for _, event := range events {
		if event.ID == eventID {
			if event.RemainingTickets > BookedTickets {
				return nil
			} else if event.RemainingTickets > 0 {
				fmt.Println("RemainingTickets is ", event.RemainingTickets)
				return errors.New("remainingTickets is less then booking ticketsg")

			} else {
				return errors.New("remainingTickets is 0")
			}
		}
	}

	return errors.New("event not found")
}

func CreateBooking(eventID1 string, BookedTickets1 string, customerName string, customerPhone string) {
	eventID, err := strconv.Atoi(eventID1)
	if err != nil {
		log.Fatal("Couldnt convert Event id to int")
	}
	BookedTickets, err := strconv.Atoi(BookedTickets1)
	if err != nil {
		log.Fatal("Couldnt convert BookedTickets to int")
	}
	newBooking := models.Booking{
		ID:            len(booking) + 1,
		CustomerName:  customerName,
		CustomerPhone: customerPhone,
		EventID:       eventID,
		BookedTickets: BookedTickets,
	}
	booking = append(booking, newBooking)
	fmt.Println("Booking Successful")
}
