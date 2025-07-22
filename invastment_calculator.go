package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5 // in percentage
func main2() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Investment Amount:")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Println("Future Value of Investment: $", futureValue)
	fmt.Println("Future Real Value of Investment: $", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
