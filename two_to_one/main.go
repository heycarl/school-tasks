package main

import (
	"fmt"
	"strings"
)

func twoToOne(s1 string, s2 string) string {
	var outStr string = ""
	alphabet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	for _, letter := range alphabet {
		if strings.Contains(s1, string(letter)) || strings.Contains(s2, string(letter)) {
			outStr += string(letter)
		}
	}
	return outStr
}

func main() {
	fmt.Println(twoToOne("loopingisfunbutdangerous", "lessdangerousthancoding"))
}
