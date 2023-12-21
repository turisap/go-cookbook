package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func get(url string) Person {
	r, err := http.Get(url)
	if err != nil {
		log.Println("Cannot get from URL", err)
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading json data:", err)
	}
	var person Person
	json.Unmarshal(data, &person)
	return person
}

//func main() {
//	person := get("https://swapi.dev/api/people/14")
//	data, err := json.MarshalIndent(&person, "", " ")
//
//	checkError(err)
//
//	fmt.Println(string(data))
//	err = os.WriteFile("han.json", data, 0644)
//	checkError(err)
//}
