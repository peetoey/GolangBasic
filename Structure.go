package main
import "fmt"

type Books struct {
	title string 
	author string
	description string
	price float64
}
func main() {

	var BookOne Books
	BookOne.title="Go Programming"
	BookOne.author="Saksit Jantaraplin"
	BookOne.description="Hello Wohoo!"
	BookOne.price=199.99
	fmt.Println(BookOne)

	Golang:=Books{title:"Go Programming", author:"Varen Suiha", price:200}
	fmt.Println(Golang)

}