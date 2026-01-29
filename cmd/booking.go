// Package cmd is used for cli booking tickets
package cmd

import (
	"fmt"
	"log"

	"github.com/HarshithRajesh/BookIt_Go/storage"
	"github.com/spf13/cobra"
)

var bookingCmd = &cobra.Command{
	Use:   "booking",
	Short: "Booking Tickets",
	Long:  "Book the event tickets",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		eventID := args[0]
		customerName := args[1]
		customerPhone := args[2]
		BookedTickets := args[3]

		err := storage.CheckEventTickets(eventID, BookedTickets)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(eventID, customerName, customerPhone, BookedTickets)
		err = storage.CreateBooking(eventID, BookedTickets, customerName, customerPhone)
		if err != nil {
			log.Println("Booking Failed")
		}
	},
}

func init() {
	rootCmd.AddCommand(bookingCmd)
}
