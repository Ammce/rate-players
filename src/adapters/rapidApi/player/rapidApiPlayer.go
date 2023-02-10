package rapid_api_player

import (
	"fmt"

	rapidapi "github.com/Ammce/rate-players/src/adapters/rapidApi"
)

type ImportPlayerRapidApi struct {
	rapidHttpClient rapidapi.RapidApiHttpClient
}

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

func (rapidApi ImportPlayerRapidApi) ImportPlayersForTeam(teamId int8) bool {
	fmt.Println("Starting to import player")
	playerResponse := Response{}
	err := rapidApi.rapidHttpClient.GET("/players?team=541&season=2022", &playerResponse)
	if err != nil {
		fmt.Println("Error has happened hereeeeeeeee")
		return false
	}
	fmt.Println("PLAYER REZPONSE", playerResponse)
	// TODO - Marshall Body to Response Struct
	// TODO -  Transform Response struct to slice of players for bulk creation for database
	return true
}

func NewImportPlayerRapidApi(rapidHttpClient rapidapi.RapidApiHttpClient) ImportPlayerRapidApi {
	return ImportPlayerRapidApi{
		rapidHttpClient: rapidHttpClient,
	}
}
