package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	ch := make(chan Person)

	go decode(ch)

	for {
		person, ok := <-ch

		if ok {
			fmt.Printf("%#v", person)
			fmt.Println("\n" + strings.Repeat("@", 20))
		} else {
			break
		}
	}
}

func decode(c chan Person) {
	f, err := os.OpenFile("./people-stream.json", os.O_APPEND|os.O_CREATE, 0644)
	checkError(err)

	defer f.Close()

	d := json.NewDecoder(f)

	for {
		var person Person

		err = d.Decode(&person)
		if err == io.EOF {
			break
		}

		checkError(err)

		c <- person
	}

	close(c)
}
