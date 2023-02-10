package rapidapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type RapidApiHttpClient struct{}

var rapidApiKey = os.Getenv("RAPID_API_KEY")
var rapidApiBaseUrl = os.Getenv("RAPID_API_BASE_URL")
var rapidApiHost = os.Getenv("RAPID_API_HOST")

func formatUrl(url string) string {
	return fmt.Sprintf("%s%s", rapidApiBaseUrl, url)
}

func (RapidApiHttpClient) GET(url string, target interface{}) error {
	fullUrl := formatUrl(url)
	req, _ := http.NewRequest("GET", fullUrl, nil)

	req.Header.Add("X-RapidAPI-Key", rapidApiKey)
	req.Header.Add("X-RapidAPI-Host", rapidApiHost)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)

}

func NewRapidApiHttpClient() RapidApiHttpClient {
	return RapidApiHttpClient{}
}
