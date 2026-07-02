package command

import (
	"errors"
	"fmt"
	"os"
)

const REMOVE_USE string = "Use: rm file1 file2 ..."

type Remove struct {
	Files []string
}

func NewRemove(args ...string) (Remove, error) {
	if len(args) < 1 {
		return Remove{}, errors.New(REMOVE_USE)
	}

	r := Remove{
		Files: args[:],
	}

	return r, nil
}

func (r Remove) Action() (string, error) {
	for _, file := range r.Files {
		if err := removeFile(file); err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("%d arquivo(s) removido(s) com sucesso.", len(r.Files)), nil
}

func removeFile(fileName string) error {
	if fileName == "" {
		return errors.New(REMOVE_USE)
	}

	if err := os.Remove(fileName); err != nil {
		return err
	}

	return nil
}
