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
	"time"

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
		fetchResponse(args)
		// for _, link := range args {
		// 	go GetHttpResponse(link)
		// 	// fmt.Println(GetHttpResponse(link))
		// }

		// GoFetchResponseData(args)

		// fmt.Println("Hello World!")
		// reader := bufio.NewReader(os.Stdin)
		// name, _ := reader.ReadString('\n')
		// // GetHttpResponse(args[0])
		// links := strings.Split(name, '\n')
		// fmt.Println(name)
		// fmt.Println("Enter link")
		// scanner := bufio.NewScanner(os.Stdin)
		// scanner.Scan()
		// in := scanner.Text()
		// wh := strings.Split(in, " ")

		// scanner := bufio.NewScanner(os.Stdin)
		// for {
		// 	fmt.Println("Enter link: ")
		// 	var exit string
		// 	for scanner.Scan() {
		// 		var lines []string
		// 		line := scanner.Text()
		// 		if line == "" {
		// 			break
		// 		}
		// 		if line == "exit" {
		// 			exit = line
		// 			break
		// 		} else {
		// 			lines = append(lines, line)
		// 		}
		// 		GoFetchResponseData(lines)
		// 	}
		// 	if err := scanner.Err(); err != nil {
		// 		log.Println(err)
		// 	}
		// 	if exit == "exit" {
		// 		break
		// 	}
		// var lines []string
		// in := scanner.Text()
		// wh := strings.Split(in, " ")
		// fmt.Println("Result:", wh[0])
		// fmt.Println("Result:", wh[1])

		// lines = append(lines, line)
		// if len(lines) > 0 {
		// 	fmt.Println()
		// 	fmt.Println("Result:")
		// 	for _, line := range lines {
		// 		fmt.Println(line)
		// 	}
		// 	fmt.Println()
		// }
		// GoFetchResponseData(lines)
		// }

		// for _, link := range wh {
		// 	fmt.Println("I Got:", link)
		// }

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

func fetchResponse(links []string) {
	// fmt.Println("HomePage Endpoint Hit")
	var wg sync.WaitGroup

	wg.Add(len(links))
	for _, url := range links {
		go func(url string) {
			time.Sleep(1 * time.Second)
			result := GetHttpResponse(url)
			fmt.Println(result)
			wg.Done()
		}(url)
	}
	wg.Wait()
	// fmt.Println("Returning Response")
	// fmt.Fprintf(w, "Responses")
}

// func fetch(url string, wg *sync.WaitGroup) {
// 	result := GetHttpResponse(url)
// 	fmt.Println(result)
// 	wg.Done()
// 	// fmt.Println(result)
// 	// return result
// }

// //Go routinure to fetch data
// func GoFetchResponseData(links []string) {
// 	c := make(chan string)
// 	var wg sync.WaitGroup

// 	for _, link := range links {
// 		wg.Add(1) // This tells the waitgroup, that there is now 1 pending operation here
// 		go FetchResponseData(link, c, &wg)
// 	}

// 	// this function literal (also called 'anonymous function' or 'lambda expression' in other languages)
// 	// is useful because 'go' needs to prefix a function and we can save some space by not declaring a whole new function for this
// 	go func() {
// 		wg.Wait() // this blocks the goroutine until WaitGroup counter is zero
// 		close(c)  // Channels need to be closed, otherwise the below loop will go on forever
// 	}() // This calls itself

// 	// this shorthand loop is syntactic sugar for an endless loop that just waits for results to come in through the 'c' channel
// 	for response := range c {
// 		fmt.Println(response)
// 	}
// }

// func FetchResponseData(url string, c chan string, wg *sync.WaitGroup) {
// 	defer (*wg).Done()
// 	result := GetHttpResponse(url)
// 	// fmt.Println(result)
// 	c <- fmt.Sprintln(result) // pump the result into the channel
// }

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
	// wg.Done()
	return string(prettyJSON)
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
