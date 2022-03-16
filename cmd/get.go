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
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/spf13/cobra"
)

type HttpResponse struct {
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
		// fmt.Println("Hello World!")

		// GetHttpResponse(args[0])
		links := []string{
			"https://github.com/fabpot",
			"https://github.com/andrew",
			"https://github.com/taylorotwell",
			"https://github.com/egoist",
			"https://github.com/HugoGiraudel",
		}
		c := make(chan string)
		var wg sync.WaitGroup

		for _, link := range links {
			wg.Add(1) // This tells the waitgroup, that there is now 1 pending operation here
			go checkUrl(link, c, &wg)
		}

		// this function literal (also called 'anonymous function' or 'lambda expression' in other languages)
		// is useful because 'go' needs to prefix a function and we can save some space by not declaring a whole new function for this
		go func() {
			wg.Wait() // this blocks the goroutine until WaitGroup counter is zero
			close(c)  // Channels need to be closed, otherwise the below loop will go on forever
		}() // This calls itself

		// this shorthand loop is syntactic sugar for an endless loop that just waits for results to come in through the 'c' channel
		for msg := range c {
			fmt.Println(msg)
		}
	},
}

//GetHttpResponse function mainly get the response by request URL
func GetHttpResponse(requestURL string) string {
	var DefaultTransport http.RoundTripper = &http.Transport{}
	request, _ := http.NewRequest("GET", requestURL, nil)
	response, _ := DefaultTransport.RoundTrip(request)
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	result := HttpResponse{
		Url:           requestURL,
		StatusCode:    response.StatusCode,
		ContentLength: int64(binary.Size(contents)),
	}
	prettyJSON, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	// fmt.Println(string(prettyJSON))
	return string(prettyJSON)
}

func checkUrl(url string, c chan string, wg *sync.WaitGroup) {
	defer (*wg).Done()
	result := GetHttpResponse(url)
	c <- result // pump the result into the channel
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
