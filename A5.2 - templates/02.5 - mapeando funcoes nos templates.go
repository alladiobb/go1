package main

import (
	"os"
	"strings"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {

	templates := []string{
		"./html/compondoMapeando/content.html",
		"./html/compondoMapeando/footer.html",
		"./html/compondoMapeando/header.html",
	}

	templatee := template.New("content.html")
	templatee.Funcs(template.FuncMap{"ToUpper": ToUpper})
	templatee = template.Must(templatee.ParseFiles(templates...))

	error := templatee.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 20},
		{"Python", 10},
	})

	if error != nil {
		panic(error)
	}
}
