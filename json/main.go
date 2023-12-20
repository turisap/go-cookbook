package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

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

//func main() {
//	var p = Person{}
//	f, err := os.Open("skywalker.json")
//	checkError(err)
//
//	defer f.Close()
//
//	data, err := io.ReadAll(f)
//	checkError(err)
//
//	err = json.Unmarshal(data, &p)
//	checkError(err)
//
//	fmt.Printf("%v", p)
//}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func unmarshal() (person *Person) {
	f, err := os.Open("skywalker.json")
	checkError(err)

	defer f.Close()

	data, err := io.ReadAll(f)
	checkError(err)

	err = json.Unmarshal(data, &person)
	checkError(err)

	fmt.Println(person)
	return
}
