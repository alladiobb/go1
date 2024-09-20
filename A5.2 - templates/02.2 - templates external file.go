package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	template := template.Must(template.New("template.html").ParseFiles("template.html"))
	error := template.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 30},
		{"Python", 10},
	})
	if error != nil {
		panic(error)
	}
}
