package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/plain")

	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Cache-Control", "no-store")

	fmt.Fprintln(w, "Olá, mundo!")
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, `{"status":"ok"}`)
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/status", status)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
