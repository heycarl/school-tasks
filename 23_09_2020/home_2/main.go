package main

import "fmt"

func main() {
	var y int

	fmt.Print("Enter year: ")
	fmt.Scan(&y)

	if y%4 != 0 || (y%100 == 0 && y%400 != 0) {
		fmt.Print("Leap")
	} else {
		fmt.Print("Not leap")
	}
}

//TODO Ввести номер года. Определить високосный год или нет
