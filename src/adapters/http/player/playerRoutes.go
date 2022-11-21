package player_http

import (
	"github.com/gin-gonic/gin"
)

func InitPlayerRoutes(routerVersion *gin.RouterGroup) {
	router := routerVersion.Group("/players")
	router.POST("/", CreatePlayer)
}
