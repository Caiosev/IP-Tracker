package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Your current version",
	Long:    `Your current version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nversion 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
