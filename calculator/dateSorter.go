package calculator

import (
	drinkvalidator "caffecalgo/drinkValidator"
	"sort"
)


func DateSorter(caffeLogs []drinkvalidator.CaffeLogger) []drinkvalidator.CaffeLogger{

	sort.Slice(caffeLogs, func(i, j int) bool {
		return caffeLogs[i].Datetime.Before(caffeLogs[j].Datetime)
	})


	return caffeLogs
}