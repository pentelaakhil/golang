package main

import "fmt"

func main() {
	x := [4]string{"Sun", " Mon", " Tue", "Wed"}

	fmt.Println("before changing: ", x)
	updateThirdElement(&x[2])
	fmt.Println("after changing: ", x)

}
func updateThirdElement(x *string) {
	*x = "Akhil"
	fmt.Println(*x)
}
