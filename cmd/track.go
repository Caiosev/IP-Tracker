package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
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
		if len(args) == 0 {
			fmt.Println("Inform an IP address")
		}
		if len(args) > 1 {
			fmt.Println("Inform only one IP address")
		}
		fmt.Println("\nTracking IP...")
		fmt.Println("")
		fetchData(args[0])
	},
}

type Ip struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
}

func fetchData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"

	res := getData(url)

	data := Ip{}

	err := json.Unmarshal(res, &data)

	if err != nil {
		fmt.Println("Error decoding response")
	}

	if data.IP == "" {
		fmt.Println("Please provide a valid IP address")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Info", "Location", "Coords"})
	t.AppendRows([]table.Row{
		{data.Org, data.City},
		{data.Postal, data.Region, data.Location},
		{data.Timezone, data.Country},
	})
	t.AppendSeparator()
	t.Render()
}
func getData(url string) []byte {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("error: ", err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("error: ", err)
	}
	return body

}

func init() {
	rootCmd.AddCommand(trackCmd)
}
