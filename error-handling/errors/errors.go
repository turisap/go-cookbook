package main

import (
	"fmt"
)

type ParseError struct {
	Reason string
}

func (ce ParseError) Error() string {
	return fmt.Sprintf("PARSE ERROR: %s", ce.Reason)
}

func (err *ParseError) Unwrap() error {
	return err.Unwrap()
}

type OutOfRangeError struct {
	Text string
}

func (ce OutOfRangeError) Error() string {
	return fmt.Sprintf("PARSE ERROR: %s", ce.Text)
}

func (err *OutOfRangeError) Unwrap() error {
	return err.Unwrap()
}

type NoInputError struct {
}

func (ce NoInputError) Error() string {
	return fmt.Sprintf("ERROR: No input")
}

func (err *NoInputError) Unwrap() error {
	return err.Unwrap()
}
