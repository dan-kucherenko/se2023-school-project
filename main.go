package main

import (
	"fmt"
	"github.com/dan-kucherenko/se-school-project/currency_rate_getter"
)

func main() {
	time, rate, _ := currency_rate_getter.GetRateBtcToUah()
	fmt.Println("Time", time)
	fmt.Printf("Rate %f", rate)
}
