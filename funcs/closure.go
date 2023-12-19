package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

func main() {
	http.HandleFunc("/hello", logger(hello))
	http.ListenAndServe(":8000", nil)
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func logger(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		end := time.Now()
		log.Printf("%s (%v)", name, end.Sub(start))
	}
}
