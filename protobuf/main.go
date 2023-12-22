package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
	"protobuf/github.com/protocolbuffers/protobuf/examples/go/tutorialpb"
	"strings"
)

func main() {
	fmt.Println("Start execution")
	k := tutorialpb.Person{
		Name:  "Kirill",
		Age:   34,
		Phone: tutorialpb.PhoneType_PHONE_TYPE_HOME,
	}

	data, err := proto.Marshal(&k)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	if err = os.WriteFile("person.txt", data, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("End execution")

	fmt.Println("\n" + strings.Repeat("@", 20))
	fmt.Println("Read file")

	bytes, err := os.ReadFile("person.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes)

	p := &tutorialpb.Person{}
	if err = proto.Unmarshal(bytes, p); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", p)
	fmt.Println(p.GetName())
}
