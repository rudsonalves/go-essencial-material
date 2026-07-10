package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "Servidor em funcionamento.")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintln(w, "Usuário criado")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/user/create", createUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
