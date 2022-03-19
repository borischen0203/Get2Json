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
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

type GetHeadResponse struct {
	Url           string `json:"Url"`
	StatusCode    int    `json:"Status-Code"`
	ContentLength int64  `json:"Content-Length"`
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get certain properties of the HTTP responses",
	Long: `This command makes HTTP request and reports on certain properties
	of the responses it receives back.`,
	// Args: cobra.RangeArgs(1, 2),TODO: input at least one parm.
	Run: func(cmd *cobra.Command, args []string) {

		scn := bufio.NewScanner(os.Stdin)
		for {
			fmt.Println("Enter Urls:")
			var lines []string
			for scn.Scan() {
				line := scn.Text()
				if len(line) == 1 {
					// Group Separator (GS ^]): ctrl-]
					if line[0] == 'q' {
						break
					}
				}
				//TODO: if isValidInput(line){
				lines = append(lines, strings.TrimSpace(line))
				//}
			}

			if len(lines) > 0 {
				fmt.Println()
				fmt.Println("Result:")
				//-----v2-----
				// fetchResponse(lines)
				//-----v2-----

				//-----v1-----
				// GoFetchResponseData(lines)
				//-----v1-----

				fetchResponse(lines)

				// for _, line := range lines {
				// 	// fmt.Println(line)
				// 	result := GetHeadResponseService(line)
				// 	prettyJSON(result)
				// }
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

		// fetchResponse(args)
		// for _, link := range args {
		// 	go GetHttpResponse(link)
		// 	// fmt.Println(GetHttpResponse(link))
		// }

		// GoFetchResponseData(args)

		// links := []string{
		// 	"https://github.com/fabpot",
		// 	"https://github.com/andrew",
		// 	"https://github.com/taylorotwell",
		// 	"https://github.com/egoist",
		// 	"https://github.com/HugoGiraudel",
		// }
		// c := make(chan string)
		// var wg sync.WaitGroup

		// for _, link := range wh {
		// 	wg.Add(1) // This tells the waitgroup, that there is now 1 pending operation here
		// 	go checkUrl(link, c, &wg)
		// }

		// // this function literal (also called 'anonymous function' or 'lambda expression' in other languages)
		// // is useful because 'go' needs to prefix a function and we can save some space by not declaring a whole new function for this
		// go func() {
		// 	wg.Wait() // this blocks the goroutine until WaitGroup counter is zero
		// 	close(c)  // Channels need to be closed, otherwise the below loop will go on forever
		// }() // This calls itself

		// // this shorthand loop is syntactic sugar for an endless loop that just waits for results to come in through the 'c' channel
		// for msg := range c {
		// 	fmt.Println(msg)
		// }
	},
}

//-----v2-----
func fetchResponse(links []string) {
	// fmt.Println("HomePage Endpoint Hit")
	var wg sync.WaitGroup
	m := make(map[int]GetHeadResponse)

	wg.Add(len(links))
	for index, url := range links {
		go func(index int, url string) {
			// time.Sleep(1 * time.Second)
			result := GetHeadResponseService(url)
			m[index] = result
			wg.Done()
		}(index, url)
	}
	wg.Wait()
	for i := 0; i < len(links); i++ {
		fmt.Println(prettyJSON(m[i]))
	}

	// fmt.Println("Returning Response")
	// fmt.Fprintf(w, "Responses")
}

//-----v2-----

//---------------V1------------------------
//Go routinure to fetch data
func GoFetchResponseData(links []string) {
	c := make(chan GetHeadResponse)
	var wg sync.WaitGroup

	for _, link := range links {
		wg.Add(1) // This tells the waitgroup, that there is now 1 pending operation here
		go FetchResponseData(link, c, &wg)
		// fmt.Println(prettyJSON(<-c))
	}

	// this function literal (also called 'anonymous function' or 'lambda expression' in other languages)
	// is useful because 'go' needs to prefix a function and we can save some space by not declaring a whole new function for this
	go func() {
		wg.Wait() // this blocks the goroutine until WaitGroup counter is zero
		close(c)  // Channels need to be closed, otherwise the below loop will go on forever
	}() // This calls itself

	// this shorthand loop is syntactic sugar for an endless loop that just waits for results to come in through the 'c' channel
	for response := range c {
		fmt.Println(prettyJSON(response))
		// fmt.Println(response)
	}
}

func FetchResponseData(url string, c chan GetHeadResponse, wg *sync.WaitGroup) {
	defer (*wg).Done()
	result := GetHeadResponseService(url)
	// fmt.Println(result)
	// time.Sleep(3 * time.Second)
	c <- result // pump the result into the channel
}

//---------------V1------------------------
var timeout = time.Duration(2 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

//Dial: dialTimeout
//GetHeadResponseService function mainly get the response head info
func GetHeadResponseService(requestURL string) GetHeadResponse {
	// var DefaultTransport http.RoundTripper = &http.Transport{Dial: dialTimeout}
	var DefaultTransport http.RoundTripper = &http.Transport{Dial: (&net.Dialer{
		Timeout: 2 * time.Second,
	}).Dial,
		TLSHandshakeTimeout: 5 * time.Second}
	request, _ := http.NewRequest("GET", requestURL, nil)
	response, err := DefaultTransport.RoundTrip(request)
	if err != nil {
		// fmt.Printf("%s", err)
		return GetHeadResponse{requestURL, 0, 0}
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		return GetHeadResponse{requestURL, 0, 0}
	}

	//Set response
	result := GetHeadResponse{
		Url:           requestURL,
		StatusCode:    response.StatusCode,
		ContentLength: int64(binary.Size(contents)),
	}
	// prettyJSON, err := json.MarshalIndent(result, "", "   ")
	// if err != nil {
	// 	log.Fatal("Failed to generate json", err)
	// }
	// fmt.Println(string(prettyJSON))
	// wg.Done()
	return result
}

func prettyJSON(result GetHeadResponse) string {
	prettyJSON, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	return string(prettyJSON)
}

//TODO: Valid input
func isValidInput(line string) bool {

	return true
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
