package main
import "fmt"

func main() {

	doSome()
	doSomething("Hi")
	doSomething("WTF")
	addition(10, 3)
	
	result:=addition2(10, 5)*2
	fmt.Println("result*2 = ", result)

}
func doSome() {
	fmt.Println("Wohoo!")
}
func doSomething(str string) {
	fmt.Println(str)
}
func addition(a int, b int) {
	fmt.Println(a+b)
}
func addition2(a int, b int) int {
	output:=a+b
	return output
}