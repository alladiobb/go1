package main

import "fmt"

func main() {

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	s = append(s, 500)

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	s = append(s, 500)

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// for i, v := range meuArray {
	// 	fmt.Printf("O valor do índice %d é: %d \n", i, v)
	// }
	// fmt.Printf("O tipo de A é %T", b)

}
