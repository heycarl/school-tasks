package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var y float64

	fmt.Print("Enter X, Y: ")
	fmt.Scan(&x, &y)

	if y > 0 && y < math.Pow(2, x) &&  math.Pow(2, x) + math.Pow(2, y) < 1{
		fmt.Println("Ok, got it")
	} else {
		fmt.Println("nop")
	}
}