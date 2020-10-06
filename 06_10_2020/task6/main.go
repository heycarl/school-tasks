package main

import (
	"fmt"
)

func main() {
	var sum float64
	var discount float64

	fmt.Print("Enter sum: ")
	fmt.Scan(&sum)

	if sum > 1000 {
		discount = sum * 0.05
	} else if  sum > 500 {
		discount = sum * 0.03
	} else {
		discount = 0
	}
	fmt.Println("Discount is: ", discount)
}