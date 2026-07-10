package main

import (
	"fmt"
	"net/http"
)

func logginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("requisição recebida:", r.Method, r.URL.Path)

		fmt.Println("Início da requisição")

		next.ServeHTTP(w, r)

		fmt.Println("Fim da requisição")
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá, mundo!"))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /hello", logginMiddleware(http.HandlerFunc(hello)))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
