package main

import (
	"fmt"
	"net/http"

	"s07/middleware"
)

func hello(w http.ResponseWriter, r *http.Request) {
	user, ok := middleware.CurrentUser(r)
	if !ok {
		http.Error(w, "Usuário não encontrado no contexto.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(
		w,
		"Olá %s! Request ID: %s\n",
		user.Email,
		middleware.GetRequestID(r),
	)
}

func broken(w http.ResponseWriter, r *http.Request) {
	items := []string{"Olá", "mundo!"}

	w.Write([]byte(items[10]))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", hello)
	mux.HandleFunc("GET /panic", broken)

	handler := middleware.Chain(
		mux,
		middleware.RequestID,
		middleware.Logging,
		middleware.Recover,
		middleware.CORS,
		middleware.Auth,
	)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}
}
