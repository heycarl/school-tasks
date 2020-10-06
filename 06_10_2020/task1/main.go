package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Print("Enter 7*8: ")
	fmt.Scan(&a)

	if 7*8 == a {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
