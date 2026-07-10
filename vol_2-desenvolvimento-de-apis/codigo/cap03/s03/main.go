package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u User) String() string {
	return fmt.Sprintf("%s (%s)", u.Name, u.Email)
}

func bodyHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler a requisição.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Body: %s\n", string(body))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "JSON inválido.", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "User: %s\n", user)
}

func main() {
	http.HandleFunc("/body/", bodyHandler)
	http.HandleFunc("/user", userHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
