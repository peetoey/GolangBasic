package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[4] = 10
	x[3] = 8
	x[2] = 6
	x[1] = 4
	x[0] = 2
	fmt.Println(x)
	fmt.Println(x[1])

	y := [5]float64{1, 2, 3, 4, 5}
	var total float64
	for _, value := range y {
		total += value
	}
	fmt.Println(total)
	fmt.Println(total / float64(len(y))) // ค่าเฉลี่ย
}
    