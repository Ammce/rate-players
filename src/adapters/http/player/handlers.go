package player_http

import (
	"net/http"

	player_service_impl "github.com/Ammce/rate-players/src/application/player"
	"github.com/gin-gonic/gin"
)

var userService = player_service_impl.NewPlayerServiceImpl()

func CreatePlayer(c *gin.Context) {
	_, err := userService.CreatePlayer()
	if err != nil {
		// TODO - Create function to map error and return error message
		c.JSON(http.StatusOK, gin.H{
			"message": "works",
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"test": true,
	})
	return
}
