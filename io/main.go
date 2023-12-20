package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var debug *log.Logger

const url string = "http://speedtest.ftp.otenet.gr/files/test1Mb.db"

func init() {
	writer := log.Writer()
	debug = log.New(writer, "DEBUG:", log.Ldate|log.Ltime|log.Lshortfile)
}
func main() {
	//readWrite()
	//copy()
	//writeToFile()
	tempFile()
}

func readWrite() {
	r, err := http.Get(url)

	if err != nil {
		debug.Println("Cannot access url")
		log.Fatal(err)
	}

	debug.Println("Got resp")

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)

	err = os.WriteFile("rw.data", data, 0755)

	if err != nil {
		debug.Println("Cannot write to a file", err)
	}

	debug.Println("File written")
}

func copy() {
	r, err := http.Get(url)

	if err != nil {
		debug.Println("Cannot access url")
		log.Fatal(err)
	}

	defer r.Body.Close()

	file, err := os.Create("copy.data")

	if err != nil {
		debug.Println("Cannot create a file")
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	io.Copy(writer, r.Body)

	writer.Flush()
	debug.Println("Flushed writer to a file")
}

func writeToFile() {
	data := []byte("Kosikosi lalalal")

	// one-go
	//os.WriteFile("1.txt", data, 0755)

	f, err := os.Create("2.txt")

	if err != nil {
		debug.Println("Could not create file")
	}
	defer f.Close()

	bytesWritten, err := f.Write(data)

	if err != nil {
		debug.Println("Could not write data to a file")
	}
	debug.Printf("Written %d bytes", bytesWritten)
}

func tempFile() {
	// temporary DIR
	fmt.Println(os.TempDir())

	dir, err := os.MkdirTemp(os.TempDir(), "testy")

	if err != nil {
		debug.Println("Cannot create a dir", err)
	}

	debug.Printf("Temp dir created", dir)

	tmpfile, err := os.CreateTemp(dir, "mytmp_*")
	if err != nil {
		log.Println("Cannot create temp file:", err)
	}

	data := []byte("Some random stuff for the temporary file")
	_, err = tmpfile.Write(data)
	if err != nil {
		log.Println("Cannot write to temp file:", err)
	}

	err = tmpfile.Close()
	if err != nil {
		log.Println("Cannot close temp file:", err)
	}

	time.Sleep(40 * time.Second)
	defer os.RemoveAll(dir)
}
