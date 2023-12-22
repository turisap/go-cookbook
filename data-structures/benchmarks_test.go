package main

import (
	"sort"
	"testing"
)

func BenchmarkSortSlice(b *testing.B) {
	var pics = []Picture{
		{1345, "D"},
		{1933, "T"},
		{1875, "R"},
		{993, "L"},
		{2011, "K"},
	}
	f := func(i, j int) bool {
		return pics[i].year < pics[j].year
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Slice(pics, f)
	}
}
func BenchmarkSortInterface(b *testing.B) {
	var pics = []Picture{
		{1345, "D"},
		{1933, "T"},
		{1875, "R"},
		{993, "L"},
		{2011, "K"},
	}
	for i := 0; i < b.N; i++ {
		sort.Sort(Gallery(pics))
	}
}
