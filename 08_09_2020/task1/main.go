package main

import (
	"fmt"
	"math"
)

func main() {
	var l float64
	fmt.Println("Enter l: ")
	fmt.Scan(&l)
	r := l / 2 * math.Pi
	s := math.Pi * math.Pow(r, 2)
	fmt.Println("Square: ")
	fmt.Println(s)
}
