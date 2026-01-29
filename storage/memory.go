// Package storage path implements bussiness logics
package storage

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/HarshithRajesh/BookIt_Go/db"
	"github.com/HarshithRajesh/BookIt_Go/models"
	"gorm.io/gorm"
)

func GetEvents() []string {
	var names []string
	db.DB.Model(&models.Event{}).Pluck("event_name", &names)

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

	var event models.Event
	res := db.DB.First(&event, eventID)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("event not found in the database")
		}
	}
	if event.RemainingTickets >= BookedTickets {
		return nil
	} else if event.RemainingTickets > 0 {
		return fmt.Errorf("only %d tickets remaining, cannot book %d", event.RemainingTickets, BookedTickets)
	} else {
		return errors.New("this event is fully booked (0 tickets remaining)")
	}
}

func CreateBooking(eventID1 string, BookedTickets1 string, customerName string, customerPhone string) error {
	eventID, err := strconv.Atoi(eventID1)
	if err != nil {
		log.Fatal("Couldnt convert Event id to int")
	}
	BookedTickets, err := strconv.Atoi(BookedTickets1)
	if err != nil {
		log.Fatal("Couldnt convert BookedTickets to int")
	}
	newBooking := models.Booking{
		CustomerName:  customerName,
		CustomerPhone: customerPhone,
		EventID:       eventID,
		BookedTickets: BookedTickets,
	}
	result := db.DB.Create(&newBooking)
	if result.Error != nil {
		return fmt.Errorf("failed to save booking: %w", result.Error)
	}
	err = db.DB.Model(&models.Event{}).
		Where("id = ?", eventID).
		Update("remaining_tickets", gorm.Expr("remaining_tickets - ?", BookedTickets)).
		Error
	if err != nil {
		return fmt.Errorf("booking saved, but failed to update event capacity: %w", err)
	}

	fmt.Println("Booking Successful!")
	return nil
}

func GetBookingInfo(bookingID int) (models.Booking, error) {
	var booking models.Booking

	result := db.DB.First(&booking, bookingID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return booking, errors.New("booking not found")
		}
		return booking, result.Error
	}

	return booking, nil
}
