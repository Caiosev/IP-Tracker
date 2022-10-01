package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// trackCmd represents the track command
var trackCmd = &cobra.Command{
	Use:     "track",
	Aliases: []string{"t"},
	Short:   "Track the specified ip",
	Long: `Run the track command followed by the specified 
	ip and to strart the trace; Example: 
		track 8.8.8.4.4`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("track called")
		if len(args) == 0 {
			fmt.Println("informe um ip")
		}
		if len(args) > 1 {
			fmt.Println("informe apenas um ip")
		}
		fmt.Println("Verificando ip")
	},
}

func init() {
	rootCmd.AddCommand(trackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
