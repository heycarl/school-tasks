package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var y float64
	var t float64

	fmt.Print("Enter X, Y: ")
	fmt.Scan(&x, &y)

	if x > 0 && y < 0 {
		t = math.Sqrt(x) + math.Sqrt(-y)
		fmt.Println(t)
	}
}