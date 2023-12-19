module test

go 1.21.5

require (
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/mux/174 v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux/180 v0.0.0-00010101000000-000000000000
	github.com/yeqown/go-qrcode v1.5.10
)

require (
	github.com/fogleman/gg v1.3.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/yeqown/reedsolomon v1.0.0 // indirect
	golang.org/x/image v0.14.0 // indirect
)

replace github.com/gorilla/mux/180 => github.com/gorilla/mux v1.8.0

replace github.com/gorilla/mux/174 => github.com/gorilla/mux v1.7.4
