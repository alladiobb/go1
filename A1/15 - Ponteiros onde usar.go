package main

func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	// Memória -> Endereço -> Valor
	//Variável (Aponta)-> Ponteiro que tem o endereço na memória
	minhaVar := 10
	minhaVar2 := 20

	println(soma(&minhaVar, &minhaVar2))
	println(minhaVar)
	println(minhaVar2)
	println(&minhaVar)
	println(&minhaVar2)

}
