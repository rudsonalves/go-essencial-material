package command

type Command interface {
	Action() (string, error)
}

func Execute(c Command) (string, error) {
	return c.Action()
}
