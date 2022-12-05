package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Test(testStr string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Executing middleware ", testStr)
		c.Next()
	}
}
