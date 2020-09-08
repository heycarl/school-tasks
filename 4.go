package main

import (
	"fmt"
)

func main () {
	var a int
	fmt.Println("Enter a: ")
	fmt.Scan(&a)
	hard := (a % 10) + (a / 10) % 10 + (a / 100)
	fmt.Println("Сумма цифр: ", hard)
}	

 