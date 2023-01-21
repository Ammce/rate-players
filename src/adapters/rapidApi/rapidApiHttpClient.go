package rapidapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RapidApiHttpClient struct{}

var rapidApiKey = os.Getenv("RAPID_API_KEY")
var rapidApiBaseUrl = os.Getenv("RAPID_API_BASE_URL")
var rapidApiHost = os.Getenv("RAPID_API_HOST")

func (RapidApiHttpClient) GET(url string) {
	fullUrl := fmt.Sprintf("%s%s", rapidApiBaseUrl, url)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", rapidApiKey)
	req.Header.Add("X-RapidAPI-Host", rapidApiHost)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error getting all players")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("CEO URL J ", fullUrl)
	fmt.Println(res)
	fmt.Println(string(body))
}

func NewRapidApiHttpClient() RapidApiHttpClient {
	return RapidApiHttpClient{}
}
