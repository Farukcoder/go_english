package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err1 := getUserInput("Revenue: ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	expenses, err2 := getUserInput("Expenses: ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	taxRate, err3 := getUserInput("Tax Rate (in %): ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Error in input:", err1, err2, err3)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: $%.2f\n", ebt)
	fmt.Printf("Profit: $%.2f\n", profit)
	fmt.Printf("Profit Ratio: %.2f%%\n", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {

	result := fmt.Sprintf("EBT: $%.2f\nProfit: $%.2f\nProfit Ratio: %.2f%%", ebt, profit, ratio)

	os.WriteFile("financial_results.txt", []byte(result), 0644)
}
func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / revenue
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		fmt.Println("Invalid input. Please enter a positive number.")
		return 0, errors.New("invalid input")
	}
	return userInput, nil
}
