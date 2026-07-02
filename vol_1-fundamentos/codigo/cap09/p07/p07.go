package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Birthday() int {
	p.Age++
	return p.Age
}

func (p Person) String() string {
	str := "ano"
	if p.Age > 1 {
		str = "Anos"
	}
	return fmt.Sprintf("%s (%d %s)", p.Name, p.Age, str)
}

func main() {
	p1 := Person{"Maria", 28}
	p2 := Person{Name: "Gabriel", Age: 30}

	fmt.Println(p1)
	fmt.Println(p2)

	p1.Birthday()
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1)
	fmt.Println("Gabriel fez", p2.Birthday(), "anos")
}
