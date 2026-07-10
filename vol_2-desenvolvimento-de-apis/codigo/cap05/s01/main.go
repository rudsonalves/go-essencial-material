package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email,omitempty"`
	Active bool   `json:"active"`
}

func main() {
	user := User{12, "Pedro", "pedro@email.com", true}

	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// var user User

	jsonData := []byte(`{
		"id": 11,
		"name": "Fernanda",
		"email": "",
		"active": false
	}`)

	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
