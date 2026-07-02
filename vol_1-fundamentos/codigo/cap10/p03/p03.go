package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
	Age  int
}

func (d Dog) Speak() string {
	return "Au au..."
}

func (d Dog) String() string {
	return fmt.Sprintf("%s (%d)", d.Name, d.Age)
}

type Cat struct {
	Name string
	Age  int
}

func (c Cat) Speak() string {
	return "Miauuu..."
}

func (c Cat) String() string {
	return fmt.Sprintf("%s (%d)", c.Name, c.Age)
}

type Bird struct {
	Name string
	Age  int
}

func (b Bird) Speak() string {
	return "Canta..."
}

// Observe que todo objeto no Go implementa o método String, por
// isto não é necessário por ete método na interface Speaker.
// ----------------------------------
// func (b Bird) String() string {
// 	return fmt.Sprintf("%s (%d)", b.Name, b.Age)
// }

func printAnimal(s Speaker) {
	fmt.Printf("%s: %s\n", s, s.Speak())
}

func main() {
	d := Dog{"Tob", 5}
	c := Cat{"Fred", 8}
	b := Bird{"Chico", 12}

	printAnimal(d)
	printAnimal(c)
	printAnimal(b)
}
