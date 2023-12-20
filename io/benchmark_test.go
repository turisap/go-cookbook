package main

import "testing"

func BenchmarkReadWrite(b *testing.B) {
	readWrite()
}
func BenchmarkCopy(b *testing.B) {
	copy()
}
