package main

import (
	"fmt"
)

func main()  {
	var num string

	fmt.Print("Enter num: ")
	fmt.Scan(&num)
	fmt.Println(check(num))
}

func check(s string) bool {
	var a, b byte
	for i := 0; i < int(len(s) / 2); i++ {
		a += s[i]
		b += s[len(s)-i-1]
	}
	if a == b {
		return true
	} else {
		return false
	}
}