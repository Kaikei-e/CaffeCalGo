package main

import (
	//"net/http"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()
	router.LoadHTMLGlob("templates/*/**")
	router.Static("/assets", "./assets")

	router.GET("/", func (ctx *gin.Context){
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "CaffeCalGo",
		})
	})


	router.Run(":8081")



}


