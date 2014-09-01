package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	// TODO(morrita): Move to flag?
	port := os.Getenv("PORT")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello!")
	})

	r.Run(":" + port)
}
