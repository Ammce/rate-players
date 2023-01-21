package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Ammce/rate-players/src/adapters/database/postgres"
	"github.com/Ammce/rate-players/src/adapters/gin_rest"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db := postgres.Connect()
	r := gin.Default()
	gin_rest.InitRoutes(&r.RouterGroup, db)

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
