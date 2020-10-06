package main

import (
	"fmt"
)

func solve(s string) string {
	s, spaces := reverse(s)
	for i := 0; i < len(spaces); i++ {
		fmt.Println(s)
		s = s[:i+spaces[i]] + " " + s[i+spaces[i]:]
	}
	fmt.Println(len(s))
	return s
}

func reverse(s string) (string, []int) {
	out := make([]rune, len(s))
	var spaces []int;
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			spaces = append(spaces, i)
		} else {
			out[len(out)-i-1] = rune(s[i])
		}
	}
	return string(out), spaces
}
func main() {
	fmt.Println(solve("your code rocks"))
	// skco redo cruoy

}
