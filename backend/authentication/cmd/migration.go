package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "authentication service migration command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("authentication service migration command called")
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)
}
