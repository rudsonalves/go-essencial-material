package main

import (
	"fmt"
	"os"
	"p05/command"
)

func main() {
	if len(os.Args[:]) < 2 {
		fmt.Println("Use: comando arg1 arg2 ...")
		os.Exit(1)
	}

	cmd := os.Args[1]
	var args = []string{}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	var c command.Command
	var err error

	switch cmd {
	case "cat":
		c, err = command.NewCat(args...)
	case "cp":
		c, err = command.NewCopy(args...)
	case "rm":
		c, err = command.NewRemove(args...)
	case "mv":
		c, err = command.NewMove(args...)
	default:
		fmt.Println("Comandos disponíveis: cat, cp, mv, rm")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	msg, err := command.Execute(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if msg != "" {
		fmt.Println(msg)
	}
}
