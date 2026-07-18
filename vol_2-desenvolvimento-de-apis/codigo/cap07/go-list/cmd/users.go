package main

import "net/http"

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create user"))
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list users"))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update user"))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}
