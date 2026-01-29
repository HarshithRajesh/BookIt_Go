// Package cmd is used add menu commands
package cmd

import (
	"fmt"

	"github.com/HarshithRajesh/BookIt_Go/storage"
	"github.com/spf13/cobra"
)

var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "Menu Command",
	Run: func(cmd *cobra.Command, args []string) {
		event := storage.GetEvents()
		for _, n := range event {
			fmt.Println(n)
		}
	},
}

func init() {
	rootCmd.AddCommand(menuCmd)
}
