package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000.00, errors.New("could not read balance file, using default balance of 1000.00")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000.00, errors.New("could not parse balance, using default balance of 1000.00")
	}
	return balance, nil
}

func main3() {

	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Using default balance of 1000.00")
		accountBalance = 1000.00
		writeBalanceToFile(accountBalance)
		return
	}
	count := 200

	for i := 0; i < count; i++ {
		fmt.Println("Welcome to Go Bank!!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice (1-4): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your current balance is: %.2f\n", accountBalance)
		case 2:
			fmt.Println("Enter amount to deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than zero!")
				continue
			}

			accountBalance += depositAmount
			fmt.Printf("You have successfully deposited %.2f. New balance is: %.2f\n", depositAmount, accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Println("Enter amount to withdraw:")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds!")
			} else {
				accountBalance -= withdrawAmount
				fmt.Printf("You have successfully withdrawn %.2f. New balance is: %.2f\n", withdrawAmount, accountBalance)
			}
		case 4:
			fmt.Println("Thank you for using Go Bank! Goodbye!")
			continue
		default:
			fmt.Println("Invalid choice! Please try again.")
			continue
		}
	}

	fmt.Println("You have reached the maximum number of transactions allowed. Exiting the program.")

}
