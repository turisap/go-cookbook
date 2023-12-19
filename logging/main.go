package main

import (
	"io"
	"log"
	"os"
)

var (
	info  *log.Logger
	debug *log.Logger
)

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot get wd")
	}

	file, err := os.OpenFile(cwd+"/log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal("Cannot use log file", err)
	}

	writer := io.MultiWriter(file, os.Stdout)

	log.SetOutput(writer)

	info = log.New(os.Stderr, "INFO\t", log.LstdFlags)
	debug = log.New(os.Stderr, "DEBUG\t", log.LstdFlags)
}

func main() {

	info.Println("info test")
	debug.Println("info test")

	info.Println("ENV:", os.Getenv("ENV"))
	if os.Getenv("ENV") != "production" {
		debug.Println("Some debugging event happened")
	}
}

func levels() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Println("UTC date")
}

func exiting() {

	defer log.Println("Deffered execution")
	log.Println("Start execution")

	log.Panicf("Fatal error")
	log.Println("End execution")
}
