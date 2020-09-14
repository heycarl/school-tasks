package main

import (
	"fmt"
	"math"
)

func digitalRootForN(n int) int {
	var src = n
	if n == 0 {
		return 0
	}
	var count int = 0
	for n != 0 {
		count++
		n /= 10
	}
	var digits []int
	var tens []int
	for i := count-1; i > -1; i-- {
		tens = append(tens, int(math.Pow10(i)))
	}
	var j int = 0
	for src != 0 {
		digits = append(digits, src / tens[j])
		src %= tens[j]
		j++
	}

	result := 0  
	for _, v := range digits {  
		result += v 
	}  
	return result
}

func main() {
	var root int = digitalRootForN(132189)
	for root > 9 {
		root = digitalRootForN(root)
	}
	fmt.Println(root)
}
