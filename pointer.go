package main

import (
	"fmt"
)

// func Add(x *int, y int) {
// 	*x += y
// }

// func main() {
// 	x := 5
// 	y := 10

// 	Add(&x, y)

//		fmt.Println("Value of x after addition:", x)
//		fmt.Println("Value of y remains unchanged:", y)
//	}
func main2() {
	age := 30
	fmt.Println("Age:", age)

	var agePointer *int
	agePointer = &age
	fmt.Println("Age via pointer:", *agePointer)

	adultYears := GetAdultYears(agePointer)

	fmt.Println("Adult years:", adultYears)
}

func GetAdultYears(age *int) int {
	return *age - 18
}
