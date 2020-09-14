package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, alfa float64

	fmt.Println("Enter a, b and α: ")
	fmt.Scan(&a, &b, &alfa)

	p := 2 * (a + b)
	fmt.Println("P: ", p)
	d1 := int64(math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2) - 2*a*b*math.Cos(toRadians(alfa))))
	fmt.Println("D1: ", d1)
	s := a * b * math.Sin(toRadians(alfa))
	fmt.Println("S: ", s)
}

func toRadians(angle float64) float64 {
	return angle * (math.Pi / 180)
}
