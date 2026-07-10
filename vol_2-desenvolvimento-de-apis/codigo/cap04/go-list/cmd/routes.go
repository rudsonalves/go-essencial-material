package main

import "net/http"

func registerRoutes(mux *http.ServeMux) {
	registerUserRoutes(mux)
	registerListRoutes(mux)
	registerItemRoutes(mux)
}

func registerUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("GET /users", listUsers)
	mux.HandleFunc("GET /users/{id}", getUser)
	mux.HandleFunc("PUT /users/{id}", updateUser)
	mux.HandleFunc("DELETE /users/{id}", deleteUser)
}

func registerListRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /lists", listLists)
	mux.HandleFunc("POST /lists", createList)
	mux.HandleFunc("GET /lists/{id}", getList)
	mux.HandleFunc("PUT /lists/{id}", updateList)
	mux.HandleFunc("DELETE /lists/{id}", deleteList)
}

func registerItemRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /lists/{id}/items", createItem)
	mux.HandleFunc("GET /items/{id}", getItem)
	mux.HandleFunc("PATCH /items/{id}", updateItem)
	mux.HandleFunc("DELETE /items/{id}", deleteItem)
}
