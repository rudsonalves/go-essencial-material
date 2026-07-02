package command

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const CAT_USE string = "Use: cat file_name -n\n  onde -n enumera as linhas"

type Cat struct {
	FileName  string
	Enumerate bool
}

func NewCat(args ...string) (Cat, error) {
	if len(args) < 1 {
		return Cat{}, errors.New(CAT_USE)
	}

	enumerate := false
	for _, arg := range args[1:] {
		if arg == "-n" {
			enumerate = true
		}
	}

	c := Cat{
		FileName:  args[0],
		Enumerate: enumerate,
	}

	return c, nil
}

func (c Cat) Action() (string, error) {
	if c.FileName == "" {
		return "", errors.New(CAT_USE)
	}

	file, err := os.Open(c.FileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
		if c.Enumerate {
			fmt.Printf("%4d %s\n", lineCount, scanner.Text())
			continue
		}

		fmt.Println(scanner.Text())
	}

	return "", scanner.Err()
}
