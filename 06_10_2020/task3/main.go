package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	var s float64
	var d float64

	fmt.Print("Enter R: ")
	fmt.Scan(&r)

	s = math.Pi * math.Pow(2, r) 
	d = 2 * math.Pi * r

	var delta float64
	delta = math.Round(math.Abs(s - d))
	if delta > 0 {
		fmt.Println("S > D by", delta)
	} else if delta < 0{
		fmt.Println("S < D by", delta)
	} else if delta == 0 {
		fmt.Println("Equals")
	} else {
		fmt.Println("Error")
	}
}
