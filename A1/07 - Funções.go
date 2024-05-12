package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Print(soma(10, 20))
}

func soma(a, b int) (int, bool) {
	return a + b, false
}

func somaError(a, b int) (int, error) {

	if a+b >= 100 {
		return a + b, errors.New("Soma é maior ou igual a 100!")
	}
	return a + b, nil
}
