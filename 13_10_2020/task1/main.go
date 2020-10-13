package main

import (
	"fmt"
	"math/rand"
)

var max int

func main() {
	var n int
	fmt.Println("Enter n: ")
	fmt.Scan(&n)

	// create n-size matrix
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			a[i][j] = rand.Intn(100)
		}
	}

	if n%2 == 0 {
		n = n / 2
	} else {
		n = n/2 + 1
	}
	for s := 0; s < n; s++ {
		for _, val := range a[s][s : len(a[s])-s] {
			if val > max {
				max = val
			}
		}
	}
	fmt.Println(a)
	fmt.Println(max)

}
