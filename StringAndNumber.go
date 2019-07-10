package main
import ("fmt")
func main() {

	fNumber:=60
	fSecond:=5
	fmt.Println(fNumber+fSecond)
	fmt.Println(fNumber-fSecond)
	fmt.Println(fNumber*fSecond)
	fmt.Println(fNumber/fSecond)

	p1:="Saksit Jantaraplin"
	p2:="Tutorial"

	// Concat (ต่อ string)
	pAll:= p1+p2
	fmt.Println(pAll)
	fmt.Println(pAll[5:])


}