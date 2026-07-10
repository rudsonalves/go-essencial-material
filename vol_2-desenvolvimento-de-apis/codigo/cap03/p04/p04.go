package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
)

// Base code
type Response struct {
	Data  any          `json:"data,omitempty"`
	Meta  any          `json:"meta,omitempty"`
	Error *ErrorDetail `json:"error,omitempty"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PaginationMeta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func WriteResponse(w http.ResponseWriter, status int, response Response) error {
	body, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error ao gerar resposta.", http.StatusInternalServerError)
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

// -----------   App   -----------
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u User) String() string {
	return fmt.Sprintf("User: %s (email: %s)", u.Name, u.Email)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil || mediaType != "application/json" {
		if err := WriteError(w, http.StatusUnsupportedMediaType, "UNSUPPORTED_MEDIA_TYPE", "Content-Type inválido."); err != nil {
			log.Println(err)
		}
		return
	}

	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		if err := WriteError(w, http.StatusBadRequest, "INVALID_JSON", "JSON inválido."); err != nil {
			log.Println(err)
		}
		return
	}

	if err := WriteJSON(w, http.StatusOK, user, nil); err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/users", createUserHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
