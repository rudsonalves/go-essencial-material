package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

type statusResponseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func (w *statusResponseWriter) WriteHeader(status int) {
	if w.wroteHeader {
		return
	}

	w.status = status
	w.wroteHeader = true
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusResponseWriter) Write(data []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}

	return w.ResponseWriter.Write(data)
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &statusResponseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		log.Printf(
			"%s %s %d %s %s",
			r.Method,
			r.URL.Path,
			rw.status,
			duration,
			r.RemoteAddr,
		)
	})
}

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

	handlerHello := chain(
		http.HandlerFunc(hello),
		first,
		second,
		third,
	)

	mux.Handle("GET /hello", handlerHello)

	handler := logging(mux)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}
}
