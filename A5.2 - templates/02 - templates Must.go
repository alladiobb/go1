package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	template := template.New("CursoTemplate")
	template, error := template.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}\n")

	if error != nil {
		panic(error)
	}

	error2 := template.Execute(os.Stdout, curso)

	if error2 != nil {
		panic(error2)
	}
}
