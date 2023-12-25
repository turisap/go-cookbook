package main

import (
	"log"
	"net"
)

func main() {
	raddr, err := net.ResolveUDPAddr("udp", ":9001")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte("Hello from UDP client"))
	if err != nil {
		log.Fatal(err)
	}
}
