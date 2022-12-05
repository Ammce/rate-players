package gin_rest

import (
	repository_impl_player "github.com/Ammce/rate-players/src/adapters/database/postgres/repositoryImpl/player"
	player_http "github.com/Ammce/rate-players/src/adapters/gin_rest/player"
	player_service_impl "github.com/Ammce/rate-players/src/application/player"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRoutes(router *gin.RouterGroup, db *sqlx.DB) {
	v1 := router.Group("/v1")

	var playerRepository = repository_impl_player.NewPlayerRepositoryImpl(db)
	var playerService = player_service_impl.NewPlayerServiceImpl(playerRepository)

	var playerHandlers = player_http.NewPlayerHanlders(playerService)

	player_http.InitPlayerRoutes(v1, playerHandlers)
}
