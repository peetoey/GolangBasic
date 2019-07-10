package main 
import "fmt"

func main() {
	
	go f(0) // go ตามด้วยชื่อ function

	var input string
	fmt.Scanln(&input)	

}

func f(n int) {

	for i := 1; i <= 10; i++ {
		fmt.Println(n, ":", i)
	}
}