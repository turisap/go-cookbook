package main

import (
	"fmt"
	"reflect"
)

func main() {
	// basic()
	//createStr()

	getMetadata()
}

func getMetadata() {

	// get metadata
	type Manager struct {
		Id         int    `my_package:"id"`
		GivenName  string `my_package:"given_name"`
		FamilyName string `my_package:"family_name"`
		Email      string `my_package:"email"`
	}

	m := Manager{
		Id:         1,
		GivenName:  "K",
		FamilyName: "SA",
		Email:      "sa@gmail.com",
	}

	p := reflect.TypeOf(m)

	for i := 0; i < p.NumField(); i++ {
		field := p.Field(i)
		fmt.Println(field.Tag.Get("my_package"))
	}
}

type NameFunc func(p *Person) string

type Person struct {
	id                  int32
	firstName, lastName string
	getName             NameFunc
	Roles               []string
}

func asian(person *Person) string {
	return fmt.Sprintf("%s %s", person.lastName, person.firstName)
}

func western(person *Person) string {
	return fmt.Sprintf("%s %s", person.firstName, person.lastName)
}

func (p *Person) HasRole(r string) bool {
	for _, v := range p.Roles {
		if v == r {
			return true
		}
	}

	return false
}

func createStr() {
	// two ways to create a struct instance
	p1 := Person{}
	p2 := new(Person)

	// prints out main.Person
	fmt.Printf("%T\n", p1)
	// prints out *main.Person
	fmt.Printf("%T\n", p2)
}

func basic() {

	p1 := Person{firstName: "Kirill", lastName: "Sh", getName: western}
	p2 := Person{firstName: "Li", lastName: "Katty", getName: asian}

	fmt.Println(p1.getName(&p1))
	fmt.Println(p2.getName(&p2))

	p3 := Person{firstName: "k", lastName: "s", Roles: []string{"admin", "go-dev"}}

	fmt.Println(p3.HasRole("fe"))
	fmt.Println(p3.HasRole("admin"))
}
