package main
import ("fmt")
func main() {

	fmt.Print("Input your number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output:=input*10
	fmt.Println("Your number *10 = ", output)

    
}
