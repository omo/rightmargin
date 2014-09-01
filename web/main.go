package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	// TODO(morrita): Move to flag?
	port := os.Getenv("PORT")
	// TODO(morrita): Root directory should be externalized.
	root := "/usr/local/web"
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, root+"/public/index.html")
	})

	r.Static("/public", root+"/public")
	r.Static("/bower_components", root+"/bower_components")

	r.Run(":" + port)
}
