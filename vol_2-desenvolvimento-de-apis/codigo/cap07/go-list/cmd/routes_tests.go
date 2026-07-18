package main

import "net/http"

func simulatePanic(w http.ResponseWriter, r *http.Request) {
	panic("falha simulada")
}
