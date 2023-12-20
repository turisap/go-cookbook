package main

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func BenchmarkStringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "The time is " + time.Now().Format(time.Kitchen) + " now."
	}
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{"The time is", time.Now().Format(time.Kitchen),
			"now."}, " ")
	}
}

func BenchmarkStringSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint("The time is ", time.Now().Format(time.Kitchen),

			" now.")

	}
}

func BenchmarkStringSprintDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint("The time is ", time.Now(), " now.")
	}
}
func BenchmarkStringSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("The time is %v now.", time.Now().Format(time.Kitchen))
	}
}
func BenchmarkStringSprintfDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("The time is %s now.", time.Now())
	}
}
func BenchmarkStringBuilderFprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		fmt.Fprint(&builder, "The time is ")
		fmt.Fprint(&builder, time.Now().Format(time.Kitchen))
		fmt.Fprint(&builder, " now.")
		_ = builder.String()
	}
}
func BenchmarkStringBuilderWriteString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.WriteString("The time is ")
		builder.WriteString(time.Now().Format(time.Kitchen))
		builder.WriteString(" now.")
		_ = builder.String()
	}
}
