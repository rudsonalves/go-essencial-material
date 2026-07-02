package command

import (
	"bufio"
	"errors"
	"os"
)

const COPY_USE string = "Use: cp file_origem file_destino"

type Copy struct {
	Source      string
	Destination string
}

func NewCopy(args ...string) (Copy, error) {
	if len(args) != 2 {
		return Copy{}, errors.New(COPY_USE)
	}

	c := Copy{
		Source:      args[0],
		Destination: args[1],
	}

	return c, nil
}

func (c Copy) Action() (string, error) {
	if c.Source == "" || c.Destination == "" {
		return "", errors.New(COPY_USE)
	}

	source, err := os.Open(c.Source)
	if err != nil {
		return "", err
	}
	defer source.Close()

	destination, err := os.OpenFile(
		c.Destination,
		os.O_CREATE|os.O_WRONLY|os.O_EXCL,
		0644,
	)
	if err != nil {
		return "", err
	}
	defer destination.Close()

	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		if _, err := destination.WriteString(scanner.Text() + "\n"); err != nil {
			os.Remove(c.Destination)
			return "", err
		}
	}

	if err := scanner.Err(); err != nil {
		os.Remove(c.Destination)
		return "", err
	}

	return "Cópia realizada com sucesso", nil
}
