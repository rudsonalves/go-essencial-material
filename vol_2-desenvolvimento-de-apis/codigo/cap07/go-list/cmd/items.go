package main

import "net/http"

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create item"))
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get item"))
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update item"))
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete item"))
}
