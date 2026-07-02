package main

import "fmt"

type Queue struct {
	Values []int
}

func (q Queue) string() string {
	return fmt.Sprint(q.Values)
}

func (q *Queue) Enqueue(value int) {
	q.Values = append(q.Values, value)
}

func (q *Queue) Dequeue() (int, bool) {
	if q.Size() == 0 {
		return 0, false
	}

	index := q.Size() - 1
	last := q.Values[index]
	q.Values = q.Values[:index]
	return last, true
}

func (q Queue) Size() int {
	return len(q.Values)
}

func main() {
	q := Queue{[]int{2, 7, 3, 9, 6}}

	fmt.Println(q)

	q.Enqueue(12)
	fmt.Println(q)

	q.Enqueue(43)
	fmt.Println(q)

	q.Enqueue(1)
	fmt.Println(q)

	fmt.Println("\nEnqueue...")
	for {
		value, ok := q.Dequeue()
		if ok {
			fmt.Println(q, ":", value)
		} else {
			fmt.Println("Vazio...")
			break
		}
	}
}
