package middleware

import (
	"go-list/internal/shared/http/response"
	"log"
	"net/http"
	"runtime/debug"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf(
					"panic recuperado: method=%s, path=%s, error=%v\n%s",
					r.Method,
					r.URL.Path,
					err,
					debug.Stack(),
				)

				httpresponse.WriteError(
					w,
					http.StatusInternalServerError,
					"Erro inesperado do servidor.",
				)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
