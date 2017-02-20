package dirContent

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
func GetDirectoryContentInJSON(dir string) []Content {
	// Retrieve all the files in the input directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var listOfContent []Content

	// Create a Content struct for each file and append it to a Content list.
	for _, file := range files {

		ext := filepath.Ext(file.Name())
		name := strings.TrimSuffix(file.Name(), ext)
		fullPath := filepath.Join(dir, file.Name())
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

	return listOfContent
}

// Delete user-chosen content from given directory.
func deleteSelectedContentFromDir(jsonSelected []byte) bool {
	var paths []Path
	jsonErr := json.Unmarshal(jsonSelected, &paths)
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
