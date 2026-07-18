package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"s07/response"
)

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

				response.WriteError(
					w,
					http.StatusInternalServerError,
					"Erro inesperado do servidor.",
				)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
