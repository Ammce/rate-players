package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Ammce/rate-players/src/adapters/database/postgres"
	player_http "github.com/Ammce/rate-players/src/adapters/http/player"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")

	player_http.InitPlayerRoutes(v1)

	// playerRoutes.Use(middlewares.Test("Amel"))

	// playerRoutes.GET("/test", func(c *gin.Context) {
	// 	abc := "abc"
	// 	abc = "env"
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": abc,
	// 	})
	// })

	// Pass this as real connection later
	postgres.Connect()

	PORT := "3000"

	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}

	fmt.Println("PORT IS", PORT)
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "works",
		})
	})
	address := fmt.Sprintf("0.0.0.0:%s", PORT)
	r.Run(address)
}
