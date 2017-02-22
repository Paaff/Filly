package content

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/paaff/Filly/error"
)

// Root directory global variable
var ROOT_DIR string

// Directory content struct
type Content struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Size int64  `json:"size"`
	Type string `json:"type"`
}

// Path struct for unmarshalling JSON array of paths.
type Path struct {
	Path string
}

// Get directory Content and return a JSON-encoded string
func GetDirectoryContentInJSON(dir string) ([]Content, *errorhandler.AppError) {

	fullDir := filepath.Join(ROOT_DIR, dir)
	// Retrieve all the files in the input directory
	files, err := ioutil.ReadDir(fullDir)
	if err != nil {
		return nil, &errorhandler.AppError{err, "Content cannot be found", 404} // Status: Not Found
	}

	// Create a Content struct for each file and append it to a Content list.
	var listOfContent []Content
	var ext string
	var name string

	for _, file := range files {
		// Determine whether it is a file or a folder
		if file.IsDir() {
			ext = "folder"
			name = file.Name()
		} else {
			ext = filepath.Ext(file.Name())
			name = strings.TrimSuffix(file.Name(), ext)
		}
		fullPath := filepath.Join(fullDir, file.Name())
		c := Content{
			Name: name,
			Path: fullPath,
			Size: file.Size(),
			Type: ext}

		listOfContent = append(listOfContent, c)
	}

	return listOfContent, nil
}

// Delete user-chosen content from given directory.
func deleteSelectedContentFromDir(jsonSelected []byte) bool {
	var paths []Path
	jsonErr := json.Unmarshal(jsonSelected, &paths) //TODO: Use Decoder
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// Run through the paths and delete them.
	for n := range paths {
		err := os.Remove(paths[n].Path)
		if err != nil {
			log.Fatal(err)
			return false
		}
	}
	return true
}
