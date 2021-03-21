package main

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {


	router := gin.Default()
	router.LoadHTMLGlob("templates/*/**")
	router.Static("/assets", "./assets")


	router.GET("/", func(ctx *gin.Context){
		drinkNumber := ""
		ctx.HTML(200, "index.tmpl", gin.H{
			"number": drinkNumber,
		})
	})

	router.POST("/", func(ctx *gin.Context) {
		drinkNumber := ctx.PostForm("drinkNum")
		fmt.Printf(drinkNumber)		

		ctx.HTML(http.StatusOK, "drinks.tmpl", gin.H{
			"number": drinkNumber, 
		})
	})

	router.Run(":8081")

}
