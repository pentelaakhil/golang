package main

import "fmt"

func main() {
	Slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(Slice[0:5])
	fmt.Println(Slice[5:])
	fmt.Println(Slice[2:7])
	fmt.Println(Slice[1:6])
}
