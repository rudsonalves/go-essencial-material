package httpresponse

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorBody struct {
	Error string `json:"error"`
}

func WriteError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(errorBody{
		Error: message,
	}); err != nil {
		log.Printf("erro ao escrever resposta JSON: %v", err)
	}
}
