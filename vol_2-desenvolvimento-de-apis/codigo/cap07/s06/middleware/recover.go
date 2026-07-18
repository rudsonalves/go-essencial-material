package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
)

type errorResponse struct {
	Error string `json:"error"`
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(errorResponse{
		Error: message,
	}); err != nil {
		log.Printf("erro ao escrever resposta JSON: %v", err)
	}
}

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic recuperado: method=%s path=%s error=%v\n%s",
					r.Method,
					r.URL.Path,
					err,
					debug.Stack(),
				)

				writeError(
					w,
					http.StatusInternalServerError,
					"Erro inesperado do servidor.",
				)

			}
		}()

		next.ServeHTTP(w, r)
	})
}
