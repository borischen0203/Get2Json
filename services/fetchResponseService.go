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
	"sync"
	"time"

	"github.com/borischen0203/Get2Json/dto"
)

func FetchResponseService(links []string) {
	// fmt.Println("HomePage Endpoint Hit")
	var wg sync.WaitGroup
	m := make(map[int]dto.HeadResponse)

	wg.Add(len(links))
	for index, url := range links {
		go func(index int, url string) {
			result := GetHeadResponse(url)
			m[index] = *result
			wg.Done()
		}(index, url)
	}
	wg.Wait()
	for i := 0; i < len(links); i++ {
		fmt.Println(prettyJSON(m[i]))
	}
}

func GetHeadResponse(req string) *dto.HeadResponse {
	validResult, err := url.ParseRequestURI(req)
	if err != nil {
		return &dto.HeadResponse{req, 0, 0}
	}
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

func prettyJSON(result dto.HeadResponse) string {
	prettyJSON, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	return string(prettyJSON)
}
