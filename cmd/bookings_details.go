package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/HarshithRajesh/BookIt_Go/storage"
)

var bookingsCmd = &cobra.Command{
	Use:     "bookinfo",
	Aliases: []string{"bi"},
	Short:   "Get Booking info",
	Long:    "Get Details of the all the users booking info",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		booking, err := storage.GetBookingInfo(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("--- Booking #%d ---\n", booking.ID)
		fmt.Printf("Customer: %s\n", booking.CustomerName)
		fmt.Printf("Event:    %v\n", booking.EventID)
		fmt.Printf("Tickets:  %d\n", booking.BookedTickets)
	},
}

func init() {
	rootCmd.AddCommand(bookingsCmd)
}
