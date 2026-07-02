package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func CopyData(r io.Reader, w io.Writer) error {
	_, err := io.Copy(w, r)
	return err
}

func main() {
	text := "Olá Brasil!\n"

	reader := strings.NewReader(text)
	buffer := &bytes.Buffer{}

	if err := CopyData(reader, buffer); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Conteúdo do buffer: ")
	fmt.Print(buffer.String())

	if err := CopyData(buffer, os.Stdout); err != nil {
		fmt.Println(err)
	}
}
