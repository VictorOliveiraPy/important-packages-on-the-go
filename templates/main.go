package main

import (
	"net/http"
	"os"
	"html/template"
)


type Curso struct {
	Nome string `json:"nome`
	CargaHoraria int `json:"cargaHoria`

}

type Cursos [] Curso


func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles(templates...))
		err := t.Execute(w, Cursos{
			Curso{
                Nome: "GO",
                CargaHoraria: 10,
            },
            Curso{
                Nome: "Python",
                CargaHoraria: 20,
            },
		})
		if err!= nil {
            panic(err)
        }
	})
	
	http.ListenAndServe(":8080", nil )

}


func mustTemplate(){
	curso := Curso{"GO", 40}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horaria {{.CargaHoraria}}"))
	err := t.Execute(os.Stdout, curso)
	if err!= nil {
			panic(err)
		}

}