package main
import "fmt"

func main() {

    fmt.Println(summation(1, 1, 1))
}
func summation(args...int) int {
	
	var total int // เก็บผลรวม
	for _ , n:=range args {
		total+=n
	}

	return total
}