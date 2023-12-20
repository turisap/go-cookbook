package main

import (
	"fmt"
	"reflect"
)

func main() {
	an := func(a int) {

	}

	fmt.Println(reflect.ValueOf(an).Pointer())
	fmt.Println(reflect.TypeOf(an))
}
