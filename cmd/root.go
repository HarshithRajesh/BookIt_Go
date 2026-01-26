package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "BookIt_Go",
	Short: "CLI Booking",
	Long:  "Ticket booking in golang CLI",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while executing BookIt_Go '%s'\n", err)
		os.Exit(1)
	}
}
