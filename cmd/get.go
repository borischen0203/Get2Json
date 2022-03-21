/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"bufio"
	"fmt"
	"os"

	"github.com/borischen0203/Get2Json/services"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get certain properties of the HTTP responses",
	Long: `This command makes HTTP request and reports on certain properties
	of the responses it receives back.`,
	Args: cobra.NoArgs,
	Run:  GetResponseCommand,
}

//This function mainly excute get command
func GetResponseCommand(cmd *cobra.Command, args []string) {
	scn := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Please enter URLs:")
		var lines []string
		for scn.Scan() {
			line := scn.Text()
			if len(line) == 1 {
				// enter q to into next step or exit
				if line[0] == 'q' {
					break
				}
			}
			lines = append(lines, line)

		}

		//Print result
		if len(lines) > 0 {
			fmt.Println()
			fmt.Println("Result:")
			services.FetchResponseService(lines)
			fmt.Println()
		}

		if err := scn.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		if len(lines) == 0 {
			break
		}
	}
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
