package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":9000")

	if err != nil {
		fmt.Println("no connection")
		log.Fatal(err)
	}

	n, err := conn.Write([]byte("hey hey from client"))
	_ = n

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	//addr, err := net.ResolveTCPAddr("tcp", ":9000")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//conn, err := net.DialTCP("tcp", nil, addr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer conn.Close()
	//conn.Write([]byte("Hello World from TCP Client"))
}
