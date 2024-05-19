package main

import "fmt"

func main() {

	var minhaVar interface{} = "Alladio Bonesso"
	println(minhaVar)
	println(minhaVar.(string))
	resultado, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", resultado, ok)
	println()
	resultado2 := minhaVar.(int)
	fmt.Printf("O valor de res é %v", resultado2)

}
