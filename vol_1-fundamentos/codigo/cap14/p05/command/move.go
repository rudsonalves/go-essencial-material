package command

import (
	"errors"
	"fmt"
	"os"
)

const MOVE_USE string = "Use: mv file_origem file_destino"

type Move struct {
	Source      string
	Destination string
}

func NewMove(args ...string) (Move, error) {
	if len(args) != 2 {
		return Move{}, errors.New(MOVE_USE)
	}

	m := Move{
		Source:      args[0],
		Destination: args[1],
	}

	return m, nil
}

func (m Move) Action() (string, error) {
	if m.Source == "" || m.Destination == "" {
		return "", errors.New(MOVE_USE)
	}

	if err := os.Rename(m.Source, m.Destination); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s movido para %s com sucesso", m.Source, m.Destination), nil
}
