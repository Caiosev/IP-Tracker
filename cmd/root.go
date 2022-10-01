package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func runUI() {
	fmt.Println("Rodando UI")
}

var rootCmd = &cobra.Command{
	Use:   "ip-tracker",
	Short: "Run UI version of ip-tracker",
	Long:  `Run UI version of ip-tracker`,
	Run: func(cmd *cobra.Command, args []string) {
		runUI()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
