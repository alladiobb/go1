package main

import "fmt"

type ID int

var (
	a ID      = 1
	b float32 = 1.0
)

func main() {
	var meuArray [3]int
	// var meuArray2 []int

	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(meuArray[0])

	for i, v := range meuArray {
		fmt.Printf("O valor do índice %d é: %d \n", i, v)
	}
	// fmt.Printf("O tipo de A é %T", b)

}
