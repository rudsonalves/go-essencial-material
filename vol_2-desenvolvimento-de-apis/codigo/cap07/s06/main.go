package main

import (
	"fmt"
	"net/http"

	"s06/middleware"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(
		w,
		"Olá! Request ID: %s\n",
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
	)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}
}
