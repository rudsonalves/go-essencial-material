package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)
	fmt.Println(r.Header.Get("User-Agent"))

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(data)

	fmt.Fprintln(w, "Olá mundo!")
}

func main() {
	http.HandleFunc("/", hello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
