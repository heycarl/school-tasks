package main

import "fmt"

func solution(str, ending string) bool {
	if str[len(ending)+1:] == ending {
		return true
	} else {
		return false
	}
}
func main() {
	solution("sasha", "ha")
}
