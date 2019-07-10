package main

import (
	"fmt"
)

func main() {

	fmt.Print("Input your number: ")
	var input float64
	fmt.Scanf("%f", &input)

	condition := input > 2
	// if condition {
	// 	fmt.Println("More than 2")
	// } else {
	// 	fmt.Println("Less than 2 (or equal)")
	// }

	if input < 2 || input == 2 {
		fmt.Println("Less than 2 (or equal)")
	} else {
		fmt.Println("More than 2")
	}
}
