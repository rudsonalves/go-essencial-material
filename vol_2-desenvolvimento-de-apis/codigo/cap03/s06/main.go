package main

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Teste no navegador:
// http://localhost:8080/text
func textHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintln(w, "Olá, mundo!")
}

// Teste no navegador:
// http://localhost:8080/json
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"status": "ok",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Erro ao gerar resposta.", http.StatusInternalServerError)
		return
	}
}

// Teste com curl, enviando o header Content-Type:
// curl -i -H "Content-Type: application/json" http://localhost:8080/content-type
func contentTypeHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	fmt.Fprintf(w, "Content-Type: %s\n", contentType)
}

// Teste com JSON válido:
//
//	curl -i -X POST http://localhost:8080/users \
//	  -H "Content-Type: application/json" \
//	  -d '{"name":"Maria","email":"maria@example.com"}'
//
// Teste com Content-Type inválido:
//
//	curl -i -X POST http://localhost:8080/users \
//	  -H "Content-Type: text/plain" \
//	  -d "Maria"
func userHandler(w http.ResponseWriter, r *http.Request) {
	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || mediaType != "application/json" {
		http.Error(w, "Content-Type inválido.", http.StatusUnsupportedMediaType)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "JSON inválido.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := map[string]User{
		"user": user,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Erro ao gerar resposta.", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/text", textHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/content-type", contentTypeHandler)
	http.HandleFunc("/users", userHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
