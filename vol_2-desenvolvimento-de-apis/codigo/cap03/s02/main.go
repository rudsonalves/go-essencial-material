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

func queryHandler(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()

	for k, p := range parameters {
		fmt.Fprintf(w, "%s: %s\n", k, p[0])
	}
}

func paginationHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		limit = 10
	}

	active, err := strconv.ParseBool(query.Get("active"))
	if err != nil {
		active = false
	}

	fmt.Fprintf(w, "Página: %d\nLimite: %d\nAtivo: %v", page, limit, active)
}

func main() {
	http.HandleFunc("/users/", usersHandler)
	http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/pagination", paginationHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
