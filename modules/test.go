package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", name)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func name(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/text.html")
	t.Execute(w, r.URL.EscapedPath()[1:])
}
