package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8082", nil)
}

func BuscaCepHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		responseWriter.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := request.URL.Query().Get("cep")
	if cepParam == "" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	cep, error := BuscaCep(cepParam)
	if error != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)

	// resultado, error := json.Marshal(cep)
	// if error != nil {
	// 	responseWriter.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// responseWriter.Write(resultado)

	//Essa linha abaixo vai funcionar como todas essas linhas acima comentadas
	json.NewEncoder(responseWriter).Encode(cep)
}

func BuscaCep(cep string) (*ViaCEP, error) {
	response, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}
	var c ViaCEP
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}
	return &c, nil
}
