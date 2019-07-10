package main
import ("fmt")
// Global Variables
var gVariable int = 500
func main() {
	localVariable:=40
	fmt.Println("Global = ", gVariable)
	fmt.Println("Local = ", localVariable)

	anotherFunction()
}

func anotherFunction(){
	// localVariable:=20
	fmt.Println("Global = ", gVariable)
	fmt.Println("Local = ", localVariable)
}