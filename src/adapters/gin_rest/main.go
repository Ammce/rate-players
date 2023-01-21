package gin_rest

import (
	repository_impl_player "github.com/Ammce/rate-players/src/adapters/database/postgres/repositoryImpl/player"
	import_data_player_http "github.com/Ammce/rate-players/src/adapters/gin_rest/importData/player"
	player_http "github.com/Ammce/rate-players/src/adapters/gin_rest/player"
	player_service_impl "github.com/Ammce/rate-players/src/application/player"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRoutes(router *gin.RouterGroup, db *sqlx.DB) {
	v1 := router.Group("/v1")

	// Player
	var playerRepository = repository_impl_player.NewPlayerRepositoryImpl(db)
	var playerService = player_service_impl.NewPlayerServiceImpl(playerRepository)
	var playerHandlers = player_http.NewPlayerHanlders(playerService)
	player_http.InitPlayerRoutes(v1, playerHandlers)

	// Import Data - Player
	var importDataPlayerHandlers = import_data_player_http.NewImportDataPlayerHandlers()
	import_data_player_http.InitImportDataPlayerRoutes(v1, importDataPlayerHandlers)

}
