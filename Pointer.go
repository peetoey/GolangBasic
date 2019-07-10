package main
import "fmt"

func main() {
	
	x:=10
	fmt.Printf("Value is %d\n", x)
	fmt.Printf("Address of x variable is %x\n", &x)

	var p *int // เก็บ address ของตัวแปร y
	p=&x
	fmt.Printf("Address of pointer variable is %x\n", p)

	*p=20
	fmt.Printf("New value is %d\n", x)
}