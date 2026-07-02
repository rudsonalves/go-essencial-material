package main

import (
	"fmt"

	"practice-tests/brdocs"
)

func main() {
	fmt.Println(brdocs.IsCPF("529.982.247-25"))
	fmt.Println(brdocs.IsCNPJ("04.252.011/0001-10"))
}
