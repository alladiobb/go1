package main

import "fmt"

type Cliente struct {
	nome string
}

func (cliente Cliente) andou() {
	cliente.nome = "Alladio"
	fmt.Printf("O cliente %v andou!", cliente.nome)

}

func main() {
	alladio := Cliente{
		nome: "Alladio",
	}
	alladio.andou()
	println()
}
