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
	"strings"

	"github.com/borischen0203/Get2Json/services"
	"github.com/spf13/cobra"
)

// type GetHeadResponse struct {
// 	Url           string `json:"Url"`
// 	StatusCode    int    `json:"Status-Code"`
// 	ContentLength int64  `json:"Content-Length"`
// }

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get certain properties of the HTTP responses",
	Long: `This command makes HTTP request and reports on certain properties
	of the responses it receives back.`,
	Args: cobra.NoArgs,
	Run:  GetResponseCommand,
	// Run: func(cmd *cobra.Command, args []string) {

	// 	scn := bufio.NewScanner(os.Stdin)
	// 	for {
	// 		fmt.Println("Please enter URLs:")
	// 		var lines []string
	// 		for scn.Scan() {
	// 			line := scn.Text()
	// 			if len(line) == 1 {
	// 				// enter q to into next step or exit
	// 				if line[0] == 'q' {
	// 					break
	// 				}
	// 			}
	// 			//Remove space before add to lines
	// 			lines = append(lines, strings.TrimSpace(line))

	// 		}

	// 		//Print result
	// 		if len(lines) > 0 {
	// 			fmt.Println()
	// 			fmt.Println("Result:")
	// 			fetchResponse(lines)
	// 			fmt.Println()
	// 		}

	// 		if err := scn.Err(); err != nil {
	// 			fmt.Fprintln(os.Stderr, err)
	// 			break
	// 		}
	// 		if len(lines) == 0 {
	// 			break
	// 		}
	// 	}

	// },
}

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
			//Remove space before add to lines
			lines = append(lines, strings.TrimSpace(line))

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

// func fetchResponseService(links []string) {
// 	// fmt.Println("HomePage Endpoint Hit")
// 	var wg sync.WaitGroup
// 	m := make(map[int]dto.GetHeadResponse)

// 	wg.Add(len(links))
// 	for index, url := range links {
// 		go func(index int, url string) {
// 			result := services.GetHeadResponseService(url)
// 			m[index] = *result
// 			wg.Done()
// 		}(index, url)
// 	}
// 	wg.Wait()
// 	for i := 0; i < len(links); i++ {
// 		fmt.Println(prettyJSON(m[i]))
// 	}
// }

//GetHeadResponseService function mainly get the response head info
// func GetHeadResponseService(req string) *GetHeadResponse {
// 	validResult, err := url.ParseRequestURI(req)
// 	if err != nil {
// 		return &GetHeadResponse{req, 0, 0}
// 	}
// 	requestURL := validResult.String()
// 	var DefaultTransport http.RoundTripper = &http.Transport{
// 		Dial:                (&net.Dialer{Timeout: 2 * time.Second}).Dial,
// 		TLSHandshakeTimeout: 5 * time.Second}
// 	request, _ := http.NewRequest("GET", requestURL, nil)
// 	response, err := DefaultTransport.RoundTrip(request)
// 	if err != nil {
// 		return &GetHeadResponse{requestURL, 0, 0}
// 	}
// 	defer response.Body.Close()

// 	contents, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return &GetHeadResponse{requestURL, 0, 0}
// 	}

// 	//Set response
// 	result := GetHeadResponse{
// 		Url:           requestURL,
// 		StatusCode:    response.StatusCode,
// 		ContentLength: int64(binary.Size(contents)),
// 	}
// 	return &result
// }

// func prettyJSON(result dto.GetHeadResponse) string {
// 	prettyJSON, err := json.MarshalIndent(result, "", "   ")
// 	if err != nil {
// 		log.Fatal("Failed to generate json", err)
// 	}
// 	return string(prettyJSON)
// }

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
