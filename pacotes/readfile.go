package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.WriteString("Hellow World!")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso, Tamanho: %v bytes\n", tamanho)
	f.Close()

	//Leitura
	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	//leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	os.Remove("arquivo.txt")
}

func leitura() {

}
