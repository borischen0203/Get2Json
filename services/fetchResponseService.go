package services

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/borischen0203/Get2Json/dto"
)

//This services mainly output header info as json format by inputting one or multiple URLs
func FetchResponseService(links []string) {
	var wg sync.WaitGroup
	m := make(map[int]dto.HeadResponse)

	wg.Add(len(links))
	for index, url := range links {
		go func(index int, url string) {
			//put response into a map, index as a key, headResponse ad value
			result := GetHeadResponse(url)
			m[index] = *result
			wg.Done()
		}(index, url)
	}
	wg.Wait()
	for i := 0; i < len(links); i++ {
		fmt.Println(PrettyJSON(m[i]))
	}
	writeJSON(m) // output a json file
}

//This function mainly get Http response by URL
func GetHeadResponse(reqURL string) *dto.HeadResponse {
	//Valid URL and remove space extra space
	validResult, err := url.ParseRequestURI(strings.TrimSpace(reqURL))
	if err != nil {
		return &dto.HeadResponse{reqURL, 0, 0}
	}

	//GET HTTP response
	requestURL := validResult.String()
	var DefaultTransport http.RoundTripper = &http.Transport{
		Dial:                (&net.Dialer{Timeout: 2 * time.Second}).Dial,
		TLSHandshakeTimeout: 5 * time.Second}
	request, _ := http.NewRequest("GET", requestURL, nil)
	response, err := DefaultTransport.RoundTrip(request)
	if err != nil {
		return &dto.HeadResponse{requestURL, 0, 0}
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &dto.HeadResponse{requestURL, 0, 0}
	}

	//Set response
	result := dto.HeadResponse{
		Url:           requestURL,
		StatusCode:    response.StatusCode,
		ContentLength: int64(binary.Size(contents)),
	}
	return &result
}

//This function mainly print pretty Json string
func PrettyJSON(result dto.HeadResponse) string {
	prettyJSON, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	return string(prettyJSON)
}

func writeJSON(m map[int]dto.HeadResponse) {
	result := make([]dto.HeadResponse, len(m))
	for k, v := range m {
		result[k] = v
	}
	file, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}
	_ = ioutil.WriteFile("output.json", file, 0644)
}
