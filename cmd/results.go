/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
	"github.com/srv-twry/f1-cli/cmd/models"
	"github.com/srv-twry/f1-cli/cmd/network"
)

var resultsYear string
var resultsRound string

// resultsCmd represents the results command
var resultsCmd = &cobra.Command{
	Use:   "results [flags]",
	Short: "Get the results of a F1 race.",
	Long: `Get the results of a particular F1 race . Usage:

results -y {year} -r {round}

For example:
"results -y 2017 -r 1" shows the results for 2017 Australian GP.
`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "http://ergast.com/api/f1/" + resultsYear + "/" + resultsRound + "/results"

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
			fmt.Println(race.RaceName, race.Circuit)
			fmt.Printf("\n")
			for idx, result := range race.Results {
				fmt.Println(idx+1, result.Number, result.Driver.GivenName, result.Driver.FamilyName, result.Constructor.Name, result.Laps, result.Grid, result.Points)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(resultsCmd)

	resultsCmd.Flags().StringVarP(&resultsYear, "year", "y", "", "Year (required)")
	resultsCmd.MarkFlagRequired("year")

	resultsCmd.Flags().StringVarP(&resultsRound, "round", "r", "", "Round (required)")
	resultsCmd.MarkFlagRequired("round")
}
