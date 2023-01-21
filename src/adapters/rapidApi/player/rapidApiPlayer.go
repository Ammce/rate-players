package rapid_api_player

import rapidapi "github.com/Ammce/rate-players/src/adapters/rapidApi"

type ImportPlayerRapidApi struct {
	rapidHttpClient rapidapi.RapidApiHttpClient
}

func (rapidApi ImportPlayerRapidApi) ImportPlayersForTeam(teamId int8) bool {
	rapidApi.rapidHttpClient.GET("")
	return true
}

func NewImportPlayerRapidApi(rapidHttpClient rapidapi.RapidApiHttpClient) ImportPlayerRapidApi {
	return ImportPlayerRapidApi{
		rapidHttpClient: rapidHttpClient,
	}
}
