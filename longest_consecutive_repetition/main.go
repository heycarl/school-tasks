package main
import "fmt"

type a struct {
	s string
	n int
}

type result struct {
	C rune // character
	L int  // co		unt
}

func longestRepetition(text string) result {
	if text == "" { 
		return result{} 
	}
	var current = a{s: "", n: 0}
	var max = a{s: "", n: 0}
	for _, c := range text {
		var sym = string(c)
		if current.s != sym {
			if current.n > max.n {
				max.s = current.s
				max.n = current.n
			}
			current.s = sym
			current.n = 1
		} else {
			current.n++
		}
	}
	if current.n > max.n {
		max = current
	}
	return result{[]rune(max.s)[0], max.n}
}

func main() {
	fmt.Println(longestRepetition("cbdeuuu900"))
}
