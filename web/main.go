package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello/world", headers)
	http.HandleFunc("/body", body)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8000", nil)
}

func upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Fprintf(w, "%v", fileHeader.Header)
	f, err := os.OpenFile("./uploaded/"+fileHeader.Filename,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

}

func body(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BODY:")
	body, _ := io.ReadAll(r.Body)
	fmt.Fprintf(w, "%s", body)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method : %s, Host : %s", r.Method, r.Host)
	fmt.Fprintf(w, "Path : %s, Query : %s\n", r.URL.Path, r.URL.Query())
}

func headers(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}
