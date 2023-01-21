package rapid_api_player

import rapidapi "github.com/Ammce/rate-players/src/adapters/rapidApi"

type ImportPlayerRapidApi struct {
	rapidHttpClient rapidapi.RapidApiHttpClient
}

func (rapidApi ImportPlayerRapidApi) ImportPlayersForTeam(teamId int8) bool {
	// b, _err := rapidApi.rapidHttpClient.GET("/players?team=541&season=2022")
	// TODO - Marshall Body to Response Struct
	// TODO -  Transform Response struct to slice of players for bulk creation for database
	return true
}

func NewImportPlayerRapidApi(rapidHttpClient rapidapi.RapidApiHttpClient) ImportPlayerRapidApi {
	return ImportPlayerRapidApi{
		rapidHttpClient: rapidHttpClient,
	}
}
