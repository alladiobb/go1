package main

import "fmt"

type Conta struct {
	saldo float64
}

func NovaConta() *Conta {
	return &Conta{saldo: 0}
}

func (conta *Conta) simular(valor float64) float64 {
	conta.saldo += valor
	return conta.saldo
}

func main() {
	conta := Conta{saldo: 100}
	conta.simular(200)
	fmt.Println(conta.saldo)
}
