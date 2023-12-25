package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func simpleServer() {
	listener, err := net.Listen("tcp", "localhost:9000")
	checkError(err)

	defer listener.Close()

	fmt.Println("Starting server")
	bytes := []byte{}
	for {
		conn, err := listener.Accept()
		checkError(err)

		go func(conn net.Conn) {
			time.Sleep(1 * time.Second)

			buf := make([]byte, 20)
			n, err := conn.Read(buf)
			_ = n
			checkError(err)

			fmt.Println(buf)
			bytes = append(bytes, buf...)
			fmt.Println(bytes)

			if err == io.EOF {
				fmt.Println("End of file")
				log.Fatal(err)
			}

			fmt.Println(string(bytes))

			defer conn.Close()
		}(conn)

		n, err := conn.Write([]byte("Hello from TCP server"))
		checkError(err)

		_ = n

	}
}

func advancedServer() {
	addr, err := net.ResolveTCPAddr("tcp", ":9000")
	checkError(err)

	listener, err := net.ListenTCP("tcp", addr)
	checkError(err)

	defer listener.Close()

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			buf := make([]byte, 1024)
			_, err := c.Read(buf)
			if err != nil {
				log.Fatal()
			}
			log.Print(string(buf))
			conn.Write([]byte("Hello from TCP server"))
			c.Close()
		}(conn)
	}
}

func main() {
	//simpleServer()
	advancedServer()
}
