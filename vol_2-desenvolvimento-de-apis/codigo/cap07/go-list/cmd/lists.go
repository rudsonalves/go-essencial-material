package main

import "net/http"

func listLists(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list lists"))
}

func createList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create list"))
}

func getList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get list"))
}

func updateList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update list"))
}

func deleteList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete list"))
}
