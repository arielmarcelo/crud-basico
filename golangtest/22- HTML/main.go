package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := usuario{"joão", "joao.felipe@gmail"}

		templates.ExecuteTemplate(w, "home.html", u)
	})
	fmt.Println("Executando da porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
