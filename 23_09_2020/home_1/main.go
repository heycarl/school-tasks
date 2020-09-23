package main

import (
	"fmt"
)

func main() {
	var a, b float64

	fmt.Print("Enter size in bytes: ")
	fmt.Scan(&a)

	b = a / 1024
	fmt.Printf("%.3f kBytes", b)
}

//TODO Ввести размер файла. Перевести байты в килобайты.
