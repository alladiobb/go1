package main

import "net/http"

func main() {
	//Função anonima pode ser usada na linha abaixo, 7, ao invés de usar abaixo na linha 13
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
	responseWriter.Header().Set("Content-Type", "text/plain")
	// responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte("Hello, World!"))
}
