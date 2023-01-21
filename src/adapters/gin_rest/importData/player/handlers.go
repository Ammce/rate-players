package import_data_player_http

import (
	"net/http"

	rapidapi "github.com/Ammce/rate-players/src/adapters/rapidApi"
	rapidApiPlayer "github.com/Ammce/rate-players/src/adapters/rapidApi/player"
	"github.com/gin-gonic/gin"
)

var rapidApiHttp = rapidapi.NewRapidApiHttpClient()
var rapidApiPlayerO = rapidApiPlayer.NewImportPlayerRapidApi(rapidApiHttp)

type ImportDataPlayerHandlers struct{}

func (ImportDataPlayerHandlers) ImportPlayersForClub(c *gin.Context) {

	rapidApiPlayerO.ImportPlayersForTeam(1)

	c.JSON(http.StatusOK, gin.H{
		"player": "Player",
	})
}

func NewImportDataPlayerHandlers() ImportDataPlayerHandlers {
	return ImportDataPlayerHandlers{}
}
