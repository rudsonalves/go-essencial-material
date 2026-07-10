package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 3 || parts[1] != "users" {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Identificador inválido.", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Usuário: %d\n", id)
}

func main() {
	http.HandleFunc("/users/", usersHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
