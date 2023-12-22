package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	elements []any
}

func (q *Queue) Enqueue(el any) {
	q.elements = append(q.elements, el)
}

func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	el := q.elements[0]

	q.elements = q.elements[1:]

	return el, nil
}

func (q *Queue) Peak() any {
	return q.elements[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func main() {
	q := Queue{elements: []any{1, 2, 3, 4}}

	fmt.Println(q.Peak())
	fmt.Println(q.Dequeue())
	q.Enqueue(9)
	fmt.Println(q.Dequeue())
	fmt.Println(q.elements)
}
