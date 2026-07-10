package main

import (
	"fmt"
	"net/http"
)

func createCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "username",
		Value:    "Maria",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)

	fmt.Fprintln(w, "Cookie criado.")
}

func readCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		http.Error(w, "Cookie não encontrado.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "  Value:    %s\n  Path:     %s\n  MaxAge:   %d\n  HttpOnly: %t\n  Secure:   %t\n",
		cookie.Value,
		cookie.Path,
		cookie.MaxAge,
		cookie.HttpOnly,
		cookie.Secure,
	)
}

func removeCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "username",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)

	fmt.Fprintln(w, "Cookie removido.")
}

func main() {
	http.HandleFunc("/", createCookieHandler)
	http.HandleFunc("/me", readCookieHandler)
	http.HandleFunc("/remove", removeCookieHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
