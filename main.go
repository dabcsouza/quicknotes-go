package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func main() {
	fmt.Println("Servidor rodando na porta 5000.")

  hendler := MyHandler{}
	http.ListenAndServe(":5000", hendler)
}
