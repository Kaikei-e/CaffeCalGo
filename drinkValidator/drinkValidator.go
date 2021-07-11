package drinkvalidator

import (
	//"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CaffeLogger struct{
	Number int `form:"numDrinks"`
	Method int `form:"calMethods"`
	CaffeineMg int `form:"caffeMg"`
	Amount int `form:"amount"`

}

func DrinkNum(ctx *gin.Context){
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

	logList := []CaffeLogger{}

	for i := 1; i <= numOfDrinks; i++ {
		logList = append(logList, CaffeLogger{i, 1, 0, 0})
	}

	ctx.HTML(http.StatusOK, "drinks.html", gin.H{
		"numOfDrinks": numOfDrinks,
		"logList": logList,
	})
}