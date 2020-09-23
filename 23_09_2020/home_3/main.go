package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Enter X and Y (divided by space): ")
	fmt.Scan(&x, &y)

	if x > 0 && y > 0 {
		fmt.Println("1")
	} else if x > 0 && y < 0 {
		fmt.Println("4")
	} else if x < 0 && y < 0 {
		fmt.Println("3")
	} else if x < 0 && y > 0 {
		fmt.Println("2")
	} else if x == 0 && y == 0 {
		fmt.Println("Nice joke, none of quarters")
	} else {
		fmt.Print("sd")
	}
}

//TODO Ввести координаты точки. Определить в какой координатной четверти она принадлежит?
