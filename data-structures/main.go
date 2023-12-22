package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//basics()
	//appending()
	//insertEl()
	//insertSlice()
	//removeEl()
	sorting()
}

type Person struct {
	age  uint8
	name string
}

type Picture struct {
	year uint32
	name string
}

type Gallery []Picture

func (g Gallery) Len() int {
	return len(g)
}

func (g Gallery) Less(i, j int) bool {
	return g[i].year < g[j].year
}

func (g Gallery) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func sorting() {
	integers := []int{3, 14, 159, 26, 53}
	floats := []float64{3.14, 1.41, 1.73, 2.72, 4.53}
	strs := []string{"the", "quick", "brown", "fox", "jumped"}

	sort.Ints(integers)
	sort.Float64s(floats)
	sort.Strings(strs)

	fmt.Println(integers)
	fmt.Println(floats)
	fmt.Println(strs)

	fmt.Println(strings.Repeat("@", 20) + "\n")
	fmt.Println(integers)

	// sort in reverse
	sort.Slice(integers, func(i int, j int) bool {
		return i > j
	})

	fmt.Println(integers)

	fmt.Println("\n" + strings.Repeat("@", 20))

	// sort structs
	people := []Person{
		{33, "k"},
		{28, "s"},
		{54, "n"},
		{78, "Y"},
		{13, "Y"},
	}

	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].age > people[j].age
	})
	fmt.Println(people)

	// sorting via implementing interface
	pics := []Picture{
		{1345, "D"},
		{1933, "T"},
		{1875, "R"},
		{993, "L"},
		{2011, "K"},
	}

	sort.Sort(Gallery(pics))
	fmt.Println(pics)
	//sort.Sort(sort.Reverse(Gallery(pics)))
	//fmt.Println(pics)
	//sort.Sort(sort.Reverse(Gallery(pics)))

	fmt.Println(sort.IsSorted(Gallery(pics)))
}

func insertEl() {
	idx := 3
	s := []int{0, 1, 2, 3, 4, 5}
	s = append(s[:idx+1], s[idx:]...)
	s[idx] = 999
	fmt.Println(s)
}

func insertSlice() {
	idx := 3
	s := []int{0, 1, 2, 3, 4, 5}
	inserted := []int{55, 66, 77}

	tail := append([]int{}, s[idx:]...)
	s = append(s[:idx], inserted...)
	s = append(s, tail...)
	fmt.Println(s)
}

func removeEl() {
	s := []int{1, 2, 3, 4, 5}
	idx := 3
	s = append(s[:idx], s[idx+1:]...)
	fmt.Println(s)
}

func appending() {
	//a := [3]int{0, 1, 2}
	// can it work?
	// no, you cannot append to an array
	//a1 := append(a, 3) //error
	//_ = a1
}

func basics() {

	// create slice
	s1 := make([]int, 1, 10)

	fmt.Println(s1)
	// 1
	fmt.Println(len(s1))
	// 10
	fmt.Println(cap(s1))
}
