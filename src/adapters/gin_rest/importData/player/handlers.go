package import_data_player_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImportDataPlayerHandlers struct{}

func (ImportDataPlayerHandlers) ImportPlayersForClub(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"player": "Player",
	})
}

func NewImportDataPlayerHandlers() ImportDataPlayerHandlers {
	return ImportDataPlayerHandlers{}
}
