package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Element[T constraints.Ordered] struct {
	value T
	next  *Element[T]
}
type LinkedList[T constraints.Ordered] struct {
	head *Element[T]
	size int
}

func (l *LinkedList[T]) Add(element *Element[T]) {
	if l.head == nil {
		l.head = element
	} else {
		element.next = l.head
		l.head = element
	}

	l.size++
}

func (l *LinkedList[T]) Insert(element *Element[T], marker T) {
	curr := l.head
	for curr != nil {

		if curr.value == marker {
			element.next = curr.next
			curr.next = element
		}

		curr = curr.next
	}
}

func (l *LinkedList[T]) PrintList() {
	current := l.head
	for current != nil {
		fmt.Printf("%v -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func (l *LinkedList[T]) Delete(el *Element[T]) {
	prev := l.head
	curr := l.head.next

	for curr != nil {
		if curr.value == el.value {
			prev.next = curr.next
			break
		}

		prev = curr
		curr = curr.next
	}
}

func main() {
	l := LinkedList[int]{
		head: nil,
		size: 0,
	}

	e1 := Element[int]{value: 1, next: nil}
	e2 := Element[int]{value: 2, next: nil}
	e3 := Element[int]{value: 3, next: nil}

	l.Add(&e1)
	l.Add(&e2)
	l.Add(&e3)

	l.PrintList()

	e4 := Element[int]{value: 100, next: nil}
	e5 := Element[int]{value: 200, next: nil}
	l.Insert(&e4, 2)
	l.PrintList()
	l.Insert(&e5, 100)
	l.PrintList()

	l.Delete(&e4)
	l.PrintList()
}
