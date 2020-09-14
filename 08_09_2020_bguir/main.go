package main

import (
	"fmt"
	"bufio"
	"os"
)

func main () {
	var alfa string
	fmt.Print("Enter string: ", alfa)
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	alfa = myscanner.Text()

	var sym int = 0
	var total int = 0
	var num int = 0

	fmt.Println("Enter num: ", alfa)
	fmt.Scanln(&num)

	for i := 0; i < len(alfa); i++ {
		if string(alfa[i]) == "," || string(alfa[i]) == " "{
			sym++
			if sym > num {
				total++
				sym = 0
			}
		} else {
			sym++
		}
	}
	total++ // out only
	fmt.Println("Total: ", total)
}