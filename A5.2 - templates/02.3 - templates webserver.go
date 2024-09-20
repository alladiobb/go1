package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		error := t.Execute(rw, Cursos{
			{"Go", 40},
			{"Java", 20},
			{"Python", 10},
		})

		if error != nil {
			panic(error)
		}
	})
	http.ListenAndServe(":8085", nil)

}
