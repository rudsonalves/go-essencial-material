package main

import "fmt"

type TodoList struct {
	Tasks []string
}

func (t TodoList) String() string {
	str := ""
	for index, task := range t.Tasks {
		if len(str) != 0 {
			str += fmt.Sprintf("\n%3d - %s", index+1, task)
		} else {
			str = fmt.Sprintf("%3d - %s", index+1, task)
		}
	}

	return str
}

func (t *TodoList) Add(task string) {
	t.Tasks = append(t.Tasks, task)
}

func (t TodoList) Count() int {
	return len(t.Tasks)
}

func (t *TodoList) Remove(index int) bool {
	index--
	len := t.Count()
	if index >= len {
		return false
	}

	tasks := t.Tasks[:index]
	if index < len-1 {
		tasks = append(tasks, t.Tasks[index+1:]...)
	}

	t.Tasks = tasks[:]
	return true
}

func main() {
	todo := TodoList{}

	todo.Add("Comprar bucha")
	todo.Add("Lavar o carro")
	todo.Add("Ver as notícias")

	fmt.Println(todo)
	fmt.Println(todo.Count(), "tarefas a fazer")

	ok := todo.Remove(5)
	if !ok {
		fmt.Println("Não existe tafera", 5)
	} else {
		fmt.Println(todo)
	}

	ok = todo.Remove(2)
	if !ok {
		fmt.Println("Não existe tafera", 2)
	} else {
		fmt.Println(todo)
	}
}
