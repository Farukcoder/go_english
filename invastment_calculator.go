package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount, years := 1000.0, 10.0
	expectedReturnRate := 5.5 // in percentage

	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

	futureRealValue := futureValue / math.Pow(1+expectedReturnRate/100, years)

	fmt.Println("Future Value of Investment: $", futureValue)
	fmt.Println("Future Real Value of Investment: $", futureRealValue)
}
