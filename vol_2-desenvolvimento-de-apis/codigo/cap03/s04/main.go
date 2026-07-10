package main

import (
	"fmt"
	"net/http"
)

//	curl http://localhost:8080/headers/request \
//	  -H "Authorization: Bearer abc123" \
//	  -H "Accept: application/json"
func requestHeadersHandler(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("Authorization")
	accept := r.Header.Get("Accept")
	userAgent := r.Header.Get("User-Agent")

	fmt.Fprintf(w, "Authorization: %s\n", authorization)
	fmt.Fprintf(w, "Accept: %s\n", accept)
	fmt.Fprintf(w, "User-Agent: %s\n", userAgent)
}

// curl -i http://localhost:8080/headers/response
func responseHeadersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-App-Version", "1.0")
	w.Header().Set("X-App-Id", "abc123")

	fmt.Fprintln(w, "Resposta com headers definidos pelo servidor.")
}

func main() {
	http.HandleFunc("/headers/request", requestHeadersHandler)
	http.HandleFunc("/headers/response", responseHeadersHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
