package main

import (
	"log"
	"net/http"
)

func main() {
	// http é um protocolo de comunicaçao -base da comunicaçao
	// cliente (faz requisição) - servidor (processa requisição)

	// Request - Response

	// Rotas
	// URI - Indentificador de Recurso
	//  Método -  GET, POST, PUT, DELETE

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ola Mundo"))
	})
	http.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bem vindo(a) Caroline"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))

}
