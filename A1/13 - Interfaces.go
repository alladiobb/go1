package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	//Endereco Endereco
	Endereco
}

func (e Empresa) Desativar() {
	// e.Ativo = false
}

//Uma função dentro de outra função

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	alladio := Cliente{
		Nome:  "Alladio",
		Idade: 35,
		Ativo: true,
	}

	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa)

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", alladio.Nome, alladio.Idade, alladio.Ativo)

}
