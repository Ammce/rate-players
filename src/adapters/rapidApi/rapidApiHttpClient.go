package rapidapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RapidApiHttpClient struct{}

type Player struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
}
type PlayerWithStatistics struct {
	Player Player
}

type Response struct {
	Results  int                    `json:"results"`
	Response []PlayerWithStatistics `json:"response"`
}

var rapidApiKey = os.Getenv("RAPID_API_KEY")
var rapidApiBaseUrl = os.Getenv("RAPID_API_BASE_URL")
var rapidApiHost = os.Getenv("RAPID_API_HOST")

func formatUrl(url string) string {
	return fmt.Sprintf("%s%s", rapidApiBaseUrl, url)
}

func (RapidApiHttpClient) GET(url string) ([]byte, error) {
	fullUrl := formatUrl(url)
	req, _ := http.NewRequest("GET", fullUrl, nil)

	req.Header.Add("X-RapidAPI-Key", rapidApiKey)
	req.Header.Add("X-RapidAPI-Host", rapidApiHost)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	return body, nil

}

func NewRapidApiHttpClient() RapidApiHttpClient {
	return RapidApiHttpClient{}
}
