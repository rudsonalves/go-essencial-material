package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"p06/domain"
	"p06/usecase"
	"strconv"
)

func readOption(reader *bufio.Reader) (rune, error) {
	option, _, err := reader.ReadRune()
	if err != nil {
		return 0, err
	}

	for option == '\n' || option == '\r' || option == ' ' || option == '\t' {
		option, _, err = reader.ReadRune()
		if err != nil {
			return 0, err
		}
	}

	if err := discardLine(reader); err != nil {
		return 0, err
	}

	return option, nil
}

func discardLine(reader *bufio.Reader) error {
	_, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

func inputProducts(reader *bufio.Reader, fd usecase.FileData) error {
	fmt.Println("Entra de Produtos")
	fmt.Println("Entre 'q 0' para terminar a entrada.")

	for {
		var name, strValue string

		fmt.Print("Entre produto valor: ")
		if _, err := fmt.Fscan(reader, &name, &strValue); err != nil {
			fmt.Println(err)
			continue
		}

		if name == "q" {
			break
		}

		value, err := strconv.Atoi(strValue)
		if err != nil {
			fmt.Println(err)
			continue
		}

		p := domain.NewProduct(name, value)

		if err := fd.SaveProduct(p); err != nil {
			return err
		}
	}

	return nil
}

func printProducts(fd usecase.FileData) error {
	fmt.Println("Carregando produtos")
	products, err := fd.ReadProducts()
	if err != nil {
		return err
	}

	for _, p := range products {
		fmt.Println(p)
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fd := usecase.NewFileData("data.txt")

	for {
		fmt.Println("Selecione a opção:")
		fmt.Println("n -> entrar novo produto")
		fmt.Println("p -> imprimir produtos")
		fmt.Println("q -> sair")

		option, err := readOption(reader)
		if err != nil {
			fmt.Println(err)
			return
		}

		switch option {
		case 'n':
			if err := inputProducts(reader, fd); err != nil {
				fmt.Println(err)
			}
		case 'p':
			if err := printProducts(fd); err != nil {
				fmt.Println(err)
			}
		case 'q':
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}
