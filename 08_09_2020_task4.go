package main

import (
	"fmt"
	"math"
)

func main () {
	var a float64
	fmt.Println("Enter a: ")
	fmt.Scan(&a)
	s := math.Pow(a, 2)
	s6 := s * 9
	v := math.Pow(a, 3)
	fmt.Println("S: ", s)
	fmt.Println("S all: ", s6)
	fmt.Println("V: ", v)
}