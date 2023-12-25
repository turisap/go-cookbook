package kirmath

import (
	"encoding/json"
	"io"
	"testing"
	"time"
)

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

// @COOL for runtime tests checks
func BenchmarkFibonacciWithSubBenchmark(b *testing.B) {
	testCases := []struct {
		name string
		n    int
	}{
		{"Fibonacci-1", 1},
		{"Fibonacci-5", 5},
		{"Fibonacci-10", 10},
		{"Fibonacci-20", 20},
		{"Fibonacci-30", 30},
	}
	for _, testCase := range testCases {
		testCase := testCase
		b.Run(testCase.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Fibonacci(testCase.n)
			}
		})
	}
}

// struct of unmarshaning JSON to go
type Person struct {
	Name      string    `json:"name"`
	Height    string    `json:"height"`
	Mass      string    `json:"mass"`
	HairColor string    `json:"hair_color"`
	SkinColor string    `json:"skin_color"`
	EyeColor  string    `json:"eye_color"`
	BirthYear string    `json:"birth_year"`
	Gender    string    `json:"gender"`
	Homeworld string    `json:"homeworld"`
	Films     []string  `json:"films"`
	Species   []string  `json:"species"`
	Vehicles  []string  `json:"vehicles"`
	Starships []string  `json:"starships"`
	Created   time.Time `json:"created"`
	Edited    time.Time `json:"edited"`
	URL       string    `json:"url"`
}

var jsonString string = `{"name":"Han Solo","height":"180","mass":"80",
"hair_color":"brown","skin_color":"fair","eye_color":"brown","birth_year":
"29BBY","gender":"male","homeworld":"https://swapi.dev/api/planets/22/","films":
["https://swapi.dev/api/films/1/","https://swapi.dev/api/films/2/",
"https://swapi.dev/api/films/3/"],"species":[],"vehicles":[],"starships":
["https://swapi.dev/api/starships/10/","https://swapi.dev/api/starships/22/"],
"created":"2014-12-10T16:49:14.582Z","edited":"2014-12-20T21:17:50.334Z",
"url":"https://swapi.dev/api/people/14/"}`
var jsonBytes []byte = []byte(jsonString)
var person Person

func BenchmarkWriteMarshal(b *testing.B) {
	json.Unmarshal(jsonBytes, &person)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(person)
		io.Discard.Write(data)
	}
}

func BenchmarkWriteEncoding(b *testing.B) {
	json.Unmarshal(jsonBytes, &person)
	b.ResetTimer()
	encoder := json.NewEncoder(io.Discard)
	for i := 0; i < b.N; i++ {
		encoder.Encode(person)
	}
}
