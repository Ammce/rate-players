package rapidapi

import (
	"fmt"
	"os"
)

type RapidApiHttpClient struct{}

var rapidApiKey = os.Getenv("RAPID_API_KEY")
var rapidApiBaseUrl = os.Getenv("RAPID_API_BASE_URL")
var rapidApiHost = os.Getenv("RAPID_API_HOST")

func (RapidApiHttpClient) GET(url string) {
	fullUrl := fmt.Sprintf("%s%s", rapidApiBaseUrl, url)
	fmt.Println("CEO URL J ", fullUrl)
}

func NewRapidApiHttpClient() RapidApiHttpClient {
	return RapidApiHttpClient{}
}
