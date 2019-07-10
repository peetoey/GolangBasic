package main
import "fmt"

func main() {

	x:=make(map[string] string)
	x["TH"] = "Thailand"
	x["EN"] = "England"
	x["JP"] = "Japan"
	x["H"] = "Hydrogen"
	x["Li"] = "Lithium"
	x["B"] = "Boron"
	fmt.Println(x)

	y:=map[string] string{
		"TH":"Thailand",
		"JP":"Japan",
	}
	fmt.Println(y["TH"])

}