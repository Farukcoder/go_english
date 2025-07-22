package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5 // in percentage
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Println("Investment Amount:")
	fmt.Scan(&investmentAmount)

	fmt.Println("Expected Return Rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Println("Years: ")
	fmt.Scan(&years)

	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future Value of Investment: $", futureValue)
	fmt.Println("Future Real Value of Investment: $", futureRealValue)
}
