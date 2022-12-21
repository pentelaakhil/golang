package main

import (
	"fmt"

	king "king/Queen"
)

func main() {
	var a, b int

	a = 10
	b = 20

	fmt.Println(king.Sum(a, b))
	fmt.Println(king.Sub(a, b))
	fmt.Println(king.Mul(a, b))
	fmt.Println(king.Div(a, b))
	fmt.Println(king.Divmod(a, b))

}
