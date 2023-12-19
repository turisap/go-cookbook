package main

import (
	"bytes"
	"encoding/base64"
	"github.com/gorilla/mux"
	qrcode "github.com/yeqown/go-qrcode"
	"log"
	"net/http"

	mux174 "github.com/gorilla/mux/174"
	mux180 "github.com/gorilla/mux/180"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{text}", name)
	log.Fatal(http.ListenAndServe(":8080", r))
	mux174.NewRouter()
	mux180.NewRouter()
}
func name(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	qrc, _ := qrcode.New(vars["text"])
	var buf bytes.Buffer
	qrc.SaveTo(&buf)
	base64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	_ = base64
}
