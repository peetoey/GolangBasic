package main
import "fmt"

func main() {

	for index := 0; index < 10; index++ {
		fmt.Println("Saksit Jantaraplin")
	}

	i:=1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			fmt.Println("--")
		}
		i++
	} 
	


}