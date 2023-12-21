package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", " ")

	for i := 1; i < 5; i++ {
		p := getN(i)
		encoder.Encode(p)
	}
}

type PersonJSON struct {
	Name      string    `json:"name"`
	Height    string    `json:"height,omitempty"`
	Mass      string    `json:"mass,omitempty"`
	HairColor string    `json:"hair_color,omitempty"`
	SkinColor string    `json:"skin_color,omitempty"`
	EyeColor  string    `json:"eye_color,omitempty"`
	BirthYear string    `json:"birth_year,omitempty"`
	Gender    string    `json:"gender,omitempty"`
	Homeworld string    `json:"homeworld,omitempty"`
	Films     []string  `json:"films,omitempty"`
	Species   []string  `json:"species,omitempty"`
	Vehicles  []string  `json:"vehicles,omitempty"`
	Starships []string  `json:"starships,omitempty"`
	Created   time.Time `json:"created,omitempty"`
	Edited    time.Time `json:"edited"`
	URL       string    `json:"url"`
}

func getN(n int) (person PersonJSON) {
	fmt.Println("getting person", n)
	r, err := http.Get("https://swapi.dev/api/people/" + strconv.Itoa(n))
	fmt.Println("got person", n)

	if err != nil {
		log.Println("Cannot get from URL", err)
	}
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading json data:", err)
	}
	json.Unmarshal(data, &person)
	return
}
