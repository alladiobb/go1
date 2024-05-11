package main

import "fmt"

func main() {

	salarios := map[string]int{"A": 100, "B": 200, "C": 300}
	fmt.Println(salarios["A"])

	delete(salarios, "A")

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d. \n", nome, salario)
	}

}
