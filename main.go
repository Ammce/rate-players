package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ammce/rate-players/src/adapters/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")

	playerRoutes := v1.Group("/players")

	playerRoutes.Use(middlewares.Test("Amel"))

	playerRoutes.GET("/test", func(c *gin.Context) {
		abc := "abc"
		abc = "env"
		c.JSON(http.StatusOK, gin.H{
			"message": abc,
		})
	})

	// Pass this as real connection later
	connStr := "postgres://postgres:postgres@docker.for.mac.localhost:5432/rate_players?sslmode=disable"
	_, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println("Error while connecting")
		log.Fatalln(err)
	}

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
