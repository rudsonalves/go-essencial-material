package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	registerRoutes(mux)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
