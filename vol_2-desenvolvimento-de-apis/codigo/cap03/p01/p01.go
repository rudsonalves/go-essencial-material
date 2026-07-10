package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Data  any          `json:"data,omitempty"`
	Meta  any          `json:"meta,omitempty"`
	Error *ErrorDetail `json:"error,omitempty"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PaginationMeta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func WriteResponse(w http.ResponseWriter, status int, response Response) error {
	body, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error ao gerar resposta.", http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(body)
	return err
}

func WriteJSON(w http.ResponseWriter, status int, data, meta any) error {
	return WriteResponse(w, status, Response{
		Data: data,
		Meta: meta,
	})
}

func WriteError(w http.ResponseWriter, status int, code, message string) error {
	return WriteResponse(w, status, Response{
		Error: &ErrorDetail{
			Code:    code,
			Message: message,
		},
	})
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	q := query.Get("q")

	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		limit = 10
	}

	meta := PaginationMeta{
		Page:  page,
		Limit: limit,
	}

	data := map[string]string{"q": q}

	if err = WriteJSON(w, http.StatusOK, data, meta); err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/search", searchHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
