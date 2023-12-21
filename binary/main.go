package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"time"
)

const fN = "reading_bin"

type Meter struct {
	Id        uint32
	Voltage   uint8
	Current   uint8
	Energy    uint32
	Timestamp int64
}

func main() {
	//encodeAndSave()
	decode()
}

func decode() {

	m := Meter{Id: 0, Voltage: 0, Current: 0, Energy: 0, Timestamp: 0}

	f, err := os.Open(fN)

	if err != nil {
		log.Fatal(err)
	}

	decoder := gob.NewDecoder(f)

	decoder.Decode(&m)

	fmt.Println(m)

}

func encodeAndSave() {
	m := Meter{
		Id:        1,
		Voltage:   32,
		Current:   54,
		Energy:    99,
		Timestamp: time.Now().Unix(),
	}

	f, err := os.OpenFile(fN, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	encoder := gob.NewEncoder(f)

	err = encoder.Encode(m)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File written")
}
