package dirContent

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Root directory global variable
var ROOT_DIR string = "E:\\Media"

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
func GetDirectoryContentInJSON(dir string) ([]Content, int) {

	fullDir := filepath.Join(ROOT_DIR, dir)
	var listOfContent []Content

	// Retrieve all the files in the input directory
	files, err := ioutil.ReadDir(fullDir)
	if err != nil {
		return listOfContent, 404 // Status 404 is path not available
	}

	// Create a Content struct for each file and append it to a Content list.
	for _, file := range files {

		ext := filepath.Ext(file.Name())
		name := strings.TrimSuffix(file.Name(), ext)
		fullPath := filepath.Join(fullDir, file.Name())
		c := Content{
			Name: name,
			Path: fullPath,
			Size: file.Size(),
			Type: ext}

		// Is content folder?
		if c.Type == "" {
			c.Type = "folder"
		}
		listOfContent = append(listOfContent, c)
	}

	return listOfContent, 0
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
