package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io"
	"testing"
	"time"
)

type Meter struct {
	Id        uint32
	Voltage   float32
	Current   float32
	Energy    uint32
	Timestamp uint64
}

var reading Meter = Meter{
	Id:        123456,
	Voltage:   229.5,
	Current:   1.3,
	Energy:    4321,
	Timestamp: uint64(time.Now().UnixNano()),
}
var jsonString string = `{
"Id": 123456,
"Voltage": 229.5,
"Current": 1.3,
"Energy": 4321,
"Timestamp": "2009-11-10 23:00:00 +0000 UTC"
}`

var jsonData []byte = []byte(jsonString)
var gobData *bytes.Buffer

func init() {
	gobData = bytes.NewBuffer([]byte{})
	gob.NewEncoder(gobData).Encode(reading)
}

func BenchmarkEncodeJson(b *testing.B) {
	encoder := json.NewEncoder(io.Discard)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(reading)
	}
}
func BenchmarkEncodeGob(b *testing.B) {
	encoder := gob.NewEncoder(io.Discard)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(reading)
	}
}
func BenchmarkDecodeJson(b *testing.B) {
	var data Meter
	buf := bytes.NewReader(jsonData)
	decoder := json.NewDecoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(&data)
		buf.Seek(0, 0)
	}
}
func BenchmarkDecodeGob(b *testing.B) {
	var data Meter
	buf := bytes.NewReader(gobData.Bytes())
	decoder := gob.NewDecoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(&data)
		buf.Seek(0, 0)
	}
}
