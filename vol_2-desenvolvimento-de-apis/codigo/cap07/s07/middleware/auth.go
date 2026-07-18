package middleware

import (
	"context"
	"errors"
	"net/http"
	"s07/response"
	"strings"
)

type authContextKey struct{}

var authenticatedUserKey = authContextKey{}

type AuthenticatedUser struct {
	ID    int
	Email string
}

func bearerToken(r *http.Request) (string, bool) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		return "", false
	}

	const prefix = "Bearer "

	if !strings.HasPrefix(authorization, prefix) {
		return "", false
	}

	token := strings.TrimSpace(strings.TrimPrefix(authorization, prefix))
	if token == "" {
		return "", false
	}

	return token, true
}

func findUserByToken(token string) (*AuthenticatedUser, error) {
	if token != "abc123" {
		return nil, errors.New("token inválido")
	}

	return &AuthenticatedUser{
		ID:    1,
		Email: "maria@example.com",
	}, nil
}

func CurrentUser(r *http.Request) (*AuthenticatedUser, bool) {
	user, ok := r.Context().Value(authenticatedUserKey).(*AuthenticatedUser)
	return user, ok
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := bearerToken(r)
		if !ok {
			response.WriteError(w, http.StatusUnauthorized, "Não autorizado.")
			return
		}

		user, err := findUserByToken(token)
		if err != nil {
			response.WriteError(w, http.StatusUnauthorized, "Não autorizado.")
			return
		}

		ctx := context.WithValue(r.Context(), authenticatedUserKey, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
