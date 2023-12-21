package main

import "fmt"

func main() {
	checkArrays()
}

func checkArrays() {
	var a1 [3]int

	// non-init array value
	// prints out [3]int{0, 0, 0}
	fmt.Printf("%#v", a1)
}
