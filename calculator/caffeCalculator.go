package calculator

import (
	drinkvalidator "caffecalgo/drinkValidator"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CaffeCalMethod(ctx *gin.Context){
 	caffeLogs := []drinkvalidator.CaffeLogger{}
	 
	//t := time.Time{}
	//layout := "2006/01/02 23:45"


	numOfDrinksStr := ctx.PostForm("numOfDrinks")

	numOfDrinks, err := strconv.Atoi(numOfDrinksStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < numOfDrinks; i++ {
		methodStr := ctx.PostForm("calMethods" + strconv.Itoa(i))
		methodInt, err2 := strconv.Atoi(methodStr)
		if err2 != nil {
			log.Fatal(err2)
		}

		caffeMgStr := ctx.PostForm("caffeMg" + strconv.Itoa(i))
		caffeMgInt, err3 := strconv.Atoi(caffeMgStr)
		if err3 != nil {
			log.Fatal(err3)
		}


		amountStr := ctx.PostForm("amount" + strconv.Itoa(i))
		amountInt, err4 := strconv.Atoi(amountStr)
		if err4 != nil {
			log.Fatal(err4)
		}

		datetimeStr := ctx.PostForm("datetime" + strconv.Itoa(i))
		var	caffeLogger drinkvalidator.CaffeLogger

		caffeLogger.Number = i
		caffeLogger.Method = methodInt
		caffeLogger.CaffeineMg = caffeMgInt
		caffeLogger.Amount = amountInt
		caffeLogger.Datetime = datetimeStr

		caffeLogs = append(caffeLogs, caffeLogger)

		fmt.Println(caffeLogs[i])

	}



	ctx.HTML(http.StatusOK, "calculatedPage.html", gin.H{
		
	})
}
