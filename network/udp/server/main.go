package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.ListenPacket("udp", ":9001")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		_, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Received %s from %s", string(buf), addr)
		conn.WriteTo([]byte("Hello from UDP server"), addr)
	}
}
