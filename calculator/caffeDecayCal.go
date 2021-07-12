package calculator

import (
	drinkvalidator "caffecalgo/drinkValidator"
	"log"
	"time"
)

type caffeineDecay struct{
	decayCaffe []int
	decayTime []time.Time 
}

func CaffeDecayCals(caffeLogs []drinkvalidator.CaffeLogger){
	//caffeineDecays := []caffeineDecay{}
	listLength := len(caffeLogs)
	log.Println(listLength)

	timeDuration := []time.Duration{}
	j := 0


	for i := 0; i < listLength -1; i++ {
		j += 1
		periodOfTime := caffeLogs[j].Datetime.Sub(caffeLogs[i].Datetime)

		timeDuration = append(timeDuration, periodOfTime)
	}

	log.Println(timeDuration)


	for i := 0; i < listLength - 1; i++ {
		j += 1
		if j > listLength {
			break
		}

		if i == listLength - 2{
			calMethodSimple(caffeLogs[i])
		}

		if caffeLogs[i].Method == 1 {

			calMethod1(caffeLogs[i], timeDuration[i])
		}else if (caffeLogs[i].Method == 2){
			calMethod2(caffeLogs[i], timeDuration[i])
		}
	}
}

func calMethod1(caffeStruct drinkvalidator.CaffeLogger, periodOfTime time.Duration){


}

func calMethod2(caffeStruct drinkvalidator.CaffeLogger, periodOfTime time.Duration){



}

func calMethodSimple(caffeStruct drinkvalidator.CaffeLogger){

}