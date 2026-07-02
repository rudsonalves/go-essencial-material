package main

import "fmt"

func main() {
	stock := map[string]int{
		"Arroz":    20,
		"Feijão":   8,
		"Macarrão": 5,
		"Café":     15,
	}

	for key, value := range stock {
		if value < 10 {
			fmt.Println(key)
		}
	}
}
