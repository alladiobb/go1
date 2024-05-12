package main

import (
	"fmt"
)

//Uma função dentro de outra função

func main() {

	total := func() int {
		return soma(10, 20, 30, 40, 50, 60, 70, 80, 90, 100) * 2
	}()

	fmt.Println(total)
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
