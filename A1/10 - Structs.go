package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

//Uma função dentro de outra função

func main() {

	alladio := Cliente{
		Nome:  "Alladio",
		Idade: 35,
		Ativo: true,
	}
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", alladio.Nome, alladio.Idade, alladio.Ativo)

}
