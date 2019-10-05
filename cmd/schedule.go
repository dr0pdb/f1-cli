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
	"fmt"

	"./network"

	"github.com/spf13/cobra"
)

var scheduleYear string

// scheduleCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Get the f1 schedule of the given year",
	Long: `Get the f1 schedule of the given year. For example:

go run main.go schedule 2018, will show the f1 schedule of 2018.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := network.MakeGetRequest("http://ergast.com/api/f1/" + scheduleYear)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)

	scheduleCmd.Flags().StringVarP(&scheduleYear, "year", "y", "", "Year (required)")
	scheduleCmd.MarkFlagRequired("year")
}
