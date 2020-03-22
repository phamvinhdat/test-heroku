package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context){
		c.String(http.StatusOK, "deploy-heroku success")
	})

	port := os.Getenv("PORT")
	if port == "" {
		println("port default")
		port = "8080"
	}	

	r.Run(":" + port)
}
