package main

import (
	"fmt"
)

func main() {
	fmt.Println(soma(10, 20, 30, 40, 50, 60, 70, 80, 90, 100))
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
