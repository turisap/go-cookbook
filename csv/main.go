package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type User struct {
	Id        int
	firstName string
	lastName  string
	email     string
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//regularSeparator()
	//unregularSeparator()
	writeToFile()
}

func writeToFile() {
	f, err := os.Create("./new_users.csv")
	checkError(err)

	defer f.Close()

	w := csv.NewWriter(f)

	us := [][]string{
		{"id", "firstName", "lastName", "email"},
		{"1", "Kirill", "Shakirov", "k@gmail.com"},
		{"2", "Ann", "Loly", "ako"},
	}

	/*
		err = w.WriteAll(us)
		checkError(err)
	*/

	for _, row := range us {
		err = w.Write(row)
		checkError(err)
	}

	w.Flush()
}

func regularSeparator() {
	f, err := os.Open("data.csv")

	checkError(err)

	defer f.Close()

	reader := csv.NewReader(f)

	var users []User

	// remove the header line
	reader.Read()
	rows, err := reader.ReadAll()

	checkError(err)

	for _, v := range rows {

		id, err := strconv.Atoi(v[0])

		checkError(err)

		u := User{
			Id:        id,
			firstName: v[1],
			lastName:  v[2],
			email:     v[3],
		}

		users = append(users, u)
	}

	fmt.Println(users)
}

func unregularSeparator() {
	f, err := os.Open("data-semicolon.csv")

	checkError(err)

	defer f.Close()

	reader := csv.NewReader(f)

	// custom delimiter
	reader.Comma = ';'
	// custom comment
	reader.Comment = '#'

	var users []User

	reader.Read()
	rows, err := reader.ReadAll()

	checkError(err)

	for _, v := range rows {

		id, err := strconv.Atoi(v[0])

		checkError(err)

		u := User{
			Id:        id,
			firstName: v[1],
			lastName:  v[2],
			email:     v[3],
		}

		users = append(users, u)
	}

	fmt.Println(users)
}
