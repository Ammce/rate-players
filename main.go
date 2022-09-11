package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Ammce/rate-players/src/adapters/http/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")

	playerRoutes := v1.Group("/players")

	playerRoutes.Use(middlewares.Test("Amel"))

	playerRoutes.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Yay",
		})
	})

	PORT := os.Getenv("PORT")
	fmt.Println("PORT IS", PORT)
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "works",
		})
	})
	address := fmt.Sprintf("0.0.0.0:%s", PORT)
	r.Run(address)
}
