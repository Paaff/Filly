package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/paaff/Filly/files"
)

// Root directory global variable
var ROOT_DIR string = "E:\\Media"

func main() {

	// Create a simple file server
	fs := http.FileServer(http.Dir("./web/dist"))
	http.Handle("/", fs)

	// GetDir endpoint
	http.HandleFunc("/browse", browseHandler)

	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func browseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		path := r.FormValue("path")
		if path != "" {
			// Browse from the POST form variable
			json.NewEncoder(w).Encode(dirContent.GetDirectoryContentInJSON(path))
		}

	}

}

// Set the root director.
func setRootDir(newRootDir string) {
	//TODO: Where does sanitating user input happen?
	ROOT_DIR = newRootDir
}
