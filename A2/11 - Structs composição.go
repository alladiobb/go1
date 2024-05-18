package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	//Endereco Endereco
	Endereco
}

//Uma função dentro de outra função

func main() {

	alladio := Cliente{
		Nome:  "Alladio",
		Idade: 35,
		Ativo: true,
	}

	alladio.Ativo = false
	alladio.Endereco.Cidade = "Curitiba"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", alladio.Nome, alladio.Idade, alladio.Ativo)

}
