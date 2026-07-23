package api

import (
	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine = gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Todo worker API is running",
			"status":  "succuess",
		})
	})

	router.Run(":3000")
}
