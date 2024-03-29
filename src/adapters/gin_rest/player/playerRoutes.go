package player_http

import (
	"github.com/gin-gonic/gin"
)

func InitPlayerRoutes(routerVersion *gin.RouterGroup, playerHandlers PlayerHanlders) {
	router := routerVersion.Group("/players")
	router.POST("/", playerHandlers.CreatePlayer)
	router.GET("/", playerHandlers.GetPlayers)
	router.GET("/:playerId", playerHandlers.GetPlayerById)
}
