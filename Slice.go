package main
import "fmt"

func main() {

	x:=make([]float64, 5)
	fmt.Println(x)

	slice:=[]int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2:=append(slice, 6, 7, 8, 9, 10) // ต่อ Slice
	fmt.Println(slice2)
	slice3:=make([]int, 2)
	copy(slice3, slice2)
	fmt.Println(slice3)

	arr:=[5]float64{11, 22, 33, 44, 55}
	y:=arr[0:4]
	fmt.Println(y)


}