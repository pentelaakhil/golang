package main

import "fmt"

func main() {

	//fmt.Print(" ")

	for x := 1; x <= 50; x++ {
		if x%2 != 0 {
			fmt.Print(x, "\t")
		}
	}
}
