package player_http

import (
	"fmt"
	"net/http"

	player_service_impl "github.com/Ammce/rate-players/src/application/player"
	"github.com/gin-gonic/gin"
)

var userService = player_service_impl.NewPlayerServiceImpl()

func CreatePlayer(c *gin.Context) {
	createPlayerInput := new(CreatePlayerInput)
	err := c.Bind(createPlayerInput)
	if err != nil {
		fmt.Println("Error happened while binding createPlayerInput")
		return
	}

	domainPlayer := CreatePlayerInputToPlayerEntity(*createPlayerInput)

	createdPlayer, err := userService.CreatePlayer(&domainPlayer)
	if err != nil {
		// TODO - Create function to map error and return error message
		c.JSON(http.StatusOK, gin.H{
			"message": "works",
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"player": createdPlayer,
	})
}
