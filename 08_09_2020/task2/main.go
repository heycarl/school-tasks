package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Println("Enter a, b: ")
	fmt.Scan(&a, &b)
	midar := (a + b) / 2
	fmt.Println("Среднее арифметическое: ", midar)
	midarmod := (math.Abs(a) + math.Abs(b)) / 2
	fmt.Println("Среднее арифметическое модуля: ", midarmod)
	halfdid := (a*2 - b*2) / 2
	fmt.Println("Полуразность удвоенных: ", halfdid)
	hard := math.Pow(a+b, 2) / math.Pow(a-b, 2)
	fmt.Println("Сумма в квадрате, деленное на разность в квадрате: ", hard)
}
