package calculator

import (
	drinkvalidator "caffecalgo/drinkValidator"
	"time"
)

type caffeineDecay struct{
	decayCaffe []int
	decayTime []time.Time 
}

func CaffeDecayCals(caffeLogs []drinkvalidator.CaffeLogger){
	listLength := len(caffeLogs)

	for i := 0; i < listLength; i++ {
		if caffeLogs[i].Method == 1 {
			calMethod1(caffeLogs[i])
		}else if (caffeLogs[i].Method == 2){
			calMethod2(caffeLogs[i])
		}
	}
}

func calMethod1(caffeStruct drinkvalidator.CaffeLogger){

}

func calMethod2(caffeStruct drinkvalidator.CaffeLogger){

}