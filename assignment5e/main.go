package main

import (
	"fmt"
)

func main() {

	var b string
	fmt.Println(" enter the input")
	fmt.Scan(&b)
	if b == "Golangtutorial" {
		fmt.Println("Welcome")
	} else {
		fmt.Println("End")
	}
}
