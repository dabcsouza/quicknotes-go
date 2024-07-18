package main

import (
	"fmt"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Listagem de notas</h1>")
}

func noteView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1> Visualização de nota</h1>")
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Criação de nota</h1>")
}

func main() {
	fmt.Println("Servidor rodando na porta 5000.")
	mux := http.NewServeMux()

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(":5000", mux)
}
