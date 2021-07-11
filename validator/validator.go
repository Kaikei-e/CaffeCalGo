package validator

import (
	//"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Validator(ctx *gin.Context){
	numOfDrinksStr := ctx.PostForm("numOfDrinks")

	numOfDrinks, err := strconv.Atoi(numOfDrinksStr)
	if err != nil {
		log.Fatal(err.Error())

		invalidValue := true
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"invalidValue": invalidValue,
		})
	}

	if numOfDrinks > 10 {
		invalidValue := true
		
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"invalidValue": invalidValue,
		})
	}

	

	ctx.HTML(http.StatusOK, "drinks.html", gin.H{
		"numOfDrinks": numOfDrinks,
	})
}