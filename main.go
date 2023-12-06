package main

import (
	"fmt"
	"net/http"
)

// Handler
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	// casting dari string ke []byte
	w.Write([]byte(message))
	fmt.Println("testing")
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!"
	// casting dari string ke []byte
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)

	// struct
	server := new(http.Server)
	server.Addr = address
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
