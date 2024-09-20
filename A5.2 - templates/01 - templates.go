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
	template, _ = template.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}\n")
	error := template.Execute(os.Stdout, curso)

	if error != nil {
		panic(error)
	}
}
