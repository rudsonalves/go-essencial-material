package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Olá mundo!")
}

func message(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Isto é uma HandleFunction")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/doc", message)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
