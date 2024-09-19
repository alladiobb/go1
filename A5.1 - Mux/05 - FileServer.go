package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/blog", func(rw http.ResponseWriter, r *http.Request) {
		{
			rw.Write([]byte("Hello 123!"))
		}
	})

	fileServer := http.FileServer(http.Dir("./A5.1 - http/public"))
	mux.Handle("/", fileServer)

	log.Fatal(http.ListenAndServe(":8085", mux))
}
