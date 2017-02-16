package main

import (
	"log"
	"net/http"
)

// Root directory global variable
var ROOT_DIR string = ""

func main() {

	// Create a simple file server
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

// Set the root director.
func setRootDir(newRootDir string) {
	//TODO: Where does sanitating user input happen?
	ROOT_DIR = newRootDir
}
