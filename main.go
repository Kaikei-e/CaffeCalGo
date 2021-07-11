package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static")
	
	router.GET("/", func (c *gin.Context)  {
		c.HTML(http.StatusOK, "index.html", gin.H{
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

  
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
