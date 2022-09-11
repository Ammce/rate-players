package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
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
