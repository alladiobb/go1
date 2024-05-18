package main

func main() {
	// Memória -> Endereço -> Valor
	//Variável (Aponta)-> Ponteiro que tem o endereço na memória
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20

	b := &a
	*b = 30

	println(*b)
	println(a)
	println(ponteiro)
}
