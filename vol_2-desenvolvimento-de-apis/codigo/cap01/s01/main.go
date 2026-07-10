package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Olá mundo!")
}

func main() {
	handler := HelloHandler{}

	http.ListenAndServe(":8080", handler)
}
