package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

type requestIDContextKey struct{}

var requestIDKey = requestIDContextKey{}

func newRequestID() (string, error) {
	b := make([]byte, 16)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func GetRequestID(r *http.Request) string {
	id, ok := r.Context().Value(requestIDKey).(string)
	if !ok {
		return ""
	}

	return id
}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")

		if id == "" {
			var err error

			id, err = newRequestID()
			if err != nil {
				http.Error(
					w,
					"Erro interno do servidor.",
					http.StatusInternalServerError,
				)
				return
			}
		}

		w.Header().Set("X-Request-ID", id)

		ctx := context.WithValue(r.Context(), requestIDKey, id)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
