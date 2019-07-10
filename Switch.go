package main

import "fmt"

func main() {

	fmt.Print("Input your number: ")
	var number int
	fmt.Scanf("%d", &number)

	switch number {
	case 0:
		fmt.Print("Zero")
	case 1:
		fmt.Print("One")
	case 2:
		fmt.Print("Two")
	default:
		fmt.Print("Unknown")

	}

}
