package import_data_player_http

import (
	"github.com/gin-gonic/gin"
)

func InitImportDataPlayerRoutes(routerVersion *gin.RouterGroup, importDataPlayerHandlers ImportDataPlayerHandlers) {
	router := routerVersion.Group("/import-data-player")
	router.GET("/by-team/:teamId", importDataPlayerHandlers.ImportPlayersForClub)
}
