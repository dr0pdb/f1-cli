/*
Copyright © 2019 Saurav Tiwary <srv.twry@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"github.com/srv-twry/f1-cli/cmd/models"
	"github.com/srv-twry/f1-cli/cmd/network"

	"github.com/spf13/cobra"
)

var scheduleYear string
var roundNumber string

// scheduleCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Get the f1 schedule of the given year",
	Long: `Get the f1 schedule of the given year. For example:

go run main.go schedule 2018, will show the f1 schedule of 2018.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "http://ergast.com/api/f1/" + scheduleYear

		if roundNumber != "" {
			url += "/" + roundNumber
		}

		resp, err := network.MakeGetRequest(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		data := models.Mrdata{}
		err = xml.Unmarshal(resp, &data)
		if err != nil {
			fmt.Println(err)
			return
		}

		r := data.Races
		for _, race := range r {
			fmt.Println(race.Round, race.RaceName, race.Circuit, race.Date)
		}
	},
}

func init() {
	currentYear, _, _ := time.Now().Date()
	rootCmd.AddCommand(scheduleCmd)

	scheduleCmd.Flags().StringVarP(&scheduleYear, "year", "y", strconv.Itoa(currentYear), "Year (optional)")

	scheduleCmd.Flags().StringVarP(&roundNumber, "round", "r", "", "Round (optional)")
}
