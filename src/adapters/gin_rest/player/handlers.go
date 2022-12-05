package player_http

import (
	"fmt"
	"net/http"

	"github.com/Ammce/rate-players/src/ports/player"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// TODO - CHANGE THIS TO USE SINGLE CONNECTION
// var db = postgres.GetDBInstance()
// var playerRepository = repository_impl_player.NewPlayerRepositoryImpl(db)
// var userService = player_service_impl.NewPlayerServiceImpl(playerRepository)

type PlayerHanlders struct {
	playerService player.PlayerService
}

func (ph PlayerHanlders) CreatePlayer(c *gin.Context) {
	createPlayerInput := new(CreatePlayerInput)
	err := c.Bind(createPlayerInput)
	if err != nil {
		fmt.Println("Error happened while binding createPlayerInput")
		return
	}

	domainPlayer := CreatePlayerInputToPlayerEntity(*createPlayerInput)

	createdPlayer, err := ph.playerService.CreatePlayer(&domainPlayer)
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

func NewPlayerHanlders(playerService player.PlayerService) PlayerHanlders {
	return PlayerHanlders{
		playerService: playerService,
	}
}
