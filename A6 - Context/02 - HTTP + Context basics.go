package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)

}

func handler(rw http.ResponseWriter, r *http.Request) {
	ctxt := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request Finalizada")

	select {
	case <-time.After(5 * time.Second):
		//Imprime no command line stdout
		log.Println("Request processada com sucesso!")
		//Imprime no Browser
		rw.Write([]byte("Request processada com sucesso!!!"))
	case <-ctxt.Done():
		log.Println("Request cancelada pelo cliente!!")

		http.Error(rw, "Request cancelada pelo cliente!!!!", http.StatusRequestTimeout)
	}
}
