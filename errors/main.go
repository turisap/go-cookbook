package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	g, err := processInput(os.Args)

	if err != nil {
		switch true {
		case errors.Is(NoInputError{}, err):
			fmt.Println("No input error")
			fmt.Println("No input error")
		case errors.Is(OutOfRangeError{Text: "Out of allowed range"}, err):
			fmt.Println("The number is out of range")
		case errors.Is(ParseError{Reason: "error parsing user input"}, err):
			fmt.Println("Could not  parse your intpu")
		default:
			fmt.Println("Default error")
		}
	} else {
		fmt.Println("your guess is", g)
	}
}

func processInput(arg []string) (bool, error) {
	const a = 35

	if len(arg) != 2 {
		return false, NoInputError{}
	}

	num, err := strconv.Atoi(arg[1])

	if err != nil {
		return false, ParseError{Reason: "error parsing user input"}
	}

	if num > 100 || num < 1 {
		e := OutOfRangeError{Text: "Out of allowed range"}
		return false, e
	}

	return a == num, nil
}
