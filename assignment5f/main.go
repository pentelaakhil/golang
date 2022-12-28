// Question 6
package main

import "fmt"

func main() {
	for i := 1; i <= 80; i++ {
		if i%2 == 0 && i%4 == 0 {
			fmt.Println("golang tutorial")
		} else if i%2 == 0 {
			fmt.Println("golang")
		} else if i%4 == 0 {
			fmt.Println("tutorial")

		} else if i%2 != 0 && i%4 != 0 {
			fmt.Println(i)
		}
	}
}
