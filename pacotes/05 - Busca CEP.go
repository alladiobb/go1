package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, cep := range os.Args[1:] {
		request, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if error != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", error)
		}
		defer request.Body.Close()
		response, error := io.ReadAll(request.Body)
		if error != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", error)
		}
		var data ViaCEP
		error = json.Unmarshal(response, &data)
		if error != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", error)
		}

		file, error := os.Create("cidade.txt")
		if error != nil {
			fmt.Fprintln(os.Stderr, "Erro ao criar arquivo: %v\n", error)
		}
		defer file.Close()

		_, error = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s\n", data.Cep,
			data.Localidade,
			data.Uf,
		))

		// fmt.Println(data)
	}
}
