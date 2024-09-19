package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Write([]byte("Hello World!"))
	// })
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "Blog do Alladio"})
	http.ListenAndServe(":8080", mux)

	// mux2 := http.NewServeMux()
	// mux2.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Write([]byte("Hello World2"))
	// })
	// http.ListenAndServe(":8080", mux2)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

func (blog blog) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(blog.title))
}
