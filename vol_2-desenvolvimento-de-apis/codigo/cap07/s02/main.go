package main

import (
	"fmt"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func first(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("first: antes")
		next.ServeHTTP(w, r)
		fmt.Println("first: depois")
	})
}

func second(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("second: antes")
		next.ServeHTTP(w, r)
		fmt.Println("second: depois")
	})
}

func third(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("third: antes")
		next.ServeHTTP(w, r)
		fmt.Println("third: depois")
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler")
	w.Write([]byte("Olá, mundo!"))
}

func chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}

func main() {
	mux := http.NewServeMux()

	handler := chain(
		http.HandlerFunc(hello),
		first,
		second,
		third,
	)

	mux.Handle("GET /hello", handler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
