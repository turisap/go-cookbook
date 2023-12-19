package main

import "fmt"

func A() {
	defer fmt.Println("defer on A")
	fmt.Println("A")
	B()
	fmt.Println("end of A")
}
func B() {
	//defer doNotPanic()
	defer fmt.Println("Defer on B")
	fmt.Println("B")
	//C()

	fmt.Println("end of B")
}
func C() {
	defer fmt.Println("defer on C")
	fmt.Println("C")
	panic("panic called in C")
	fmt.Println("end of C")
}
func main() {
	defer fmt.Println("defer on main")
	fmt.Println("main")
	A()
	fmt.Println("end of main")
}

func doNotPanic() {
	err := recover()
	if err != nil {
		fmt.Println("Called panic, but it is all good now")
	} else {
		fmt.Println("Defer on B")
		fmt.Println("Err", err)
	}
}
