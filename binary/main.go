package main

import (
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
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
	//decode()
	//binaryWrite()
	binaryRead()
	binaryReadManual()
}

func binaryRead() {
	var m Meter

	f, err := os.Open("binary.txt")

	if err != nil {
		log.Fatal(err)
	}

	err = binary.Read(f, binary.BigEndian, &m)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", &m)

}

func binaryReadManual() {
	var m Meter

	f, err := os.Open("binary.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	buf := make([]byte, 24)
	f.Read(buf)

	m.Id = binary.BigEndian.Uint32(buf[:4])
	m.Voltage = uint8(math.Float32frombits(binary.BigEndian.Uint32(buf[4:8])))
	m.Current = uint8(math.Float32frombits(binary.BigEndian.Uint32(buf[8:12])))
	m.Energy = binary.BigEndian.Uint32(buf[12:16])
	m.Timestamp = int64(binary.BigEndian.Uint64(buf[16:]))

	fmt.Println("\n" + strings.Repeat("@", 20))
	fmt.Printf("%v\n", m)
}

func getSimpleMeter() Meter {
	return Meter{
		Id:        1,
		Voltage:   32,
		Current:   54,
		Energy:    99,
		Timestamp: time.Now().Unix(),
	}
}

func binaryWrite() {
	f, err := os.Create("binary.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	err = binary.Write(f, binary.BigEndian, getSimpleMeter())

	if err != nil {
		log.Fatal(err)
	}
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

	defer f.Close()

	encoder := gob.NewEncoder(f)

	err = encoder.Encode(m)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File written")
}
