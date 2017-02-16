package main

import (
	"log"
	"net/http"
)

func main() {

	// Create a simple file server
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
