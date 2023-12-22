package main

import (
	"fmt"
	"testing"
)

func byCopyDown(p Person) {
	_ = fmt.Sprintf("%v", p)
}
func byReferenceDown(p *Person) {
	_ = fmt.Sprintf("%v", p)
}
func byCopyUp() Person {
	return Person{
		id:        1,
		firstName: "Sau Sheong",
		lastName:  "Chang",
	}
}
func byReferenceUp() *Person {
	return &Person{
		id:        1,
		firstName: "Sau Sheong",
		lastName:  "Chang",
	}
}

func BenchmarkStructInstanceDownCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := Person{
			id:        1,
			firstName: "Sau Sheong",
			lastName:  "Chang",
		}

		byCopyDown(p)
	}
}

func BenchmarkStructInstanceDownReference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := &Person{
			id:        1,
			firstName: "Sau Sheong",
			lastName:  "Chang",
		}
		byReferenceDown(p)
	}
}

func BenchmarkStructInstanceUpCopy(b *testing.B) {
	var p Person
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p = byCopyUp()
	}
	b.StopTimer()
	_ = fmt.Sprintf("%v", p)
}

func BenchmarkStructInstanceUpReference(b *testing.B) {
	var p *Person
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p = byReferenceUp()
	}
	b.StopTimer()
	_ = fmt.Sprintf("%v", p)
}
