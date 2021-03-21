package main

import (
	"net/http"
	"strconv"
	"fmt"
	//"err"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {


	router := gin.Default()
	router.LoadHTMLGlob("templates/*/**")
	router.Static("/assets", "./assets")


	router.GET("/", func(ctx *gin.Context){
		drinkNumber := ctx.PostForm("drinkNum")
		drinkNumber = ""
		ctx.HTML(200, "index.tmpl", gin.H{
			"number": drinkNumber,
		})
	})

	router.POST("/", func(ctx *gin.Context) {
		drinkNumber := ctx.PostForm("drinkNum")
		drinkNum, err := strconv.Atoi(drinkNumber)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(string(rune(drinkNum)))		

		var drinksRange []int
		for i := 1; i <= drinkNum; i++{
			drinksRange = append(drinksRange, i)
		}

		fmt.Println(drinksRange)
		ctx.HTML(http.StatusOK, "drinks.tmpl", gin.H{
			"number": drinksRange, 
		})
	})

	router.Run(":8081")

}
