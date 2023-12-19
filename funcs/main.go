package main

import (
	"fmt"
	"reflect"
)
import "golang.org/x/exp/constraints"

func main() {
	fmt.Println(add(2.0, 4.0))

	var a speed = 3
	var b speed = 5

	fmt.Println(add(a, b))

	fmt.Println(
		reflect.TypeOf(2.3),
	)
}

type speed int

type Number interface {
	float64 | int
}

func add[T constraints.Signed | constraints.Float](x T, y T) T {
	return x + y
}
