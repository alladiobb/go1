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

	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	fmt.Println(len(meuArray) - 1)
	// fmt.Printf("O tipo de A é %T", b)
}
