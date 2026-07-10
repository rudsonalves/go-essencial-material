package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Data  any          `json:"data,omitempty"`
	Meta  any          `json:"meta,omitempty"`
	Error *ErrorDetail `json:"error,omitempty"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PaginationMeta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func WriteResponse(w http.ResponseWriter, status int, response Response) error {
	body, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Erro ao gerar resposta.", http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(body)
	return err
}

func WriteJSON(w http.ResponseWriter, status int, data, meta any) error {
	return WriteResponse(w, status, Response{
		Data: data,
		Meta: meta,
	})
}

func WriteError(w http.ResponseWriter, status int, code, message string) error {
	return WriteResponse(w, status, Response{
		Error: &ErrorDetail{
			Code:    code,
			Message: message,
		},
	})
}

// Teste no navegador:
// http://localhost:8080/user
func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:   1,
		Name: "Maria",
	}

	if err := WriteJSON(w, http.StatusOK, user, nil); err != nil {
		log.Println(err)
		return
	}
}

// Teste no navegador:
// http://localhost:8080/users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Maria"},
		{ID: 2, Name: "João"},
	}

	meta := PaginationMeta{
		Page:  1,
		Limit: 10,
	}

	if err := WriteJSON(w, http.StatusOK, users, meta); err != nil {
		log.Println(err)
		return
	}
}

// Teste no navegador:
// http://localhost:8080/validation-error
func validationErrorHandler(w http.ResponseWriter, r *http.Request) {
	if err := WriteError(w, http.StatusBadRequest, "VALIDATION_ERROR", "Dados inválidos."); err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/validation-error", validationErrorHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
