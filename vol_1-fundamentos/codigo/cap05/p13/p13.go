package main

import "fmt"

func main() {
	capitals := map[string]int{
		"Brasília":       2982818,
		"São Paulo":      12325232,
		"Rio de Janeiro": 6211223,
		"Belo Horizonte": 2315560,
		"Salvador":       2417678,
		"Fortaleza":      2428708,
		"Recife":         1488920,
		"Curitiba":       1773718,
		"Porto Alegre":   1332570,
		"Manaus":         2063547,
		"Belém":          1303403,
		"Goiânia":        1494599,
		"Vitória":        322869,
		"Florianópolis":  537211,
		"João Pessoa":    833932,
		"Natal":          751300,
		"Maceió":         957916,
		"Aracaju":        602757,
		"Teresina":       902644,
		"São Luís":       1037775,
		"Palmas":         323625,
		"Cuiabá":         650912,
		"Campo Grande":   916001,
		"Porto Velho":    460434,
		"Rio Branco":     419452,
		"Boa Vista":      456310,
		"Macapá":         442933,
	}

	sum := 0
	for _, value := range capitals {
		sum += value
	}

	fmt.Println("Média de população das capitais brasileiras:", sum/len(capitals))
}
