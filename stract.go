package main

import (
	"fmt"
	"time"
)

type user struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

func main() {
	userFirstName := GetUserData("Enter first name:")
	userLastName := GetUserData("Enter last name:")
	userBirthDate := GetUserData("Enter birth date (YYYY-MM-DD):")

	var appUser user

	appUser = user{
		FirstName: userFirstName,
		LastName:  userLastName,
		BirthDate: userBirthDate,
		CreatedAt: time.Now(),
	}

	appUser.outputUserDetails()
}

func (u user)outputUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
