package main

import "net/http"

func main() {
	//Função anonima pode ser usada na linha abaixo, 7, ao invés de usar abaixo na linha 13
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello, World!"))
	})
	http.ListenAndServe("0:8000", nil)
}

// func BuscaCep(ResponseWriter http.ResponseWriter, r *http.Request) {
// 	ResponseWriter.Write([]byte("Hello, World!"))
// }
