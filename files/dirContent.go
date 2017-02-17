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
	Name string
	Path string
	Size int64
	Type string
}

// Path struct for unmarshalling JSON array of paths.
type Path struct {
	Path string
}

// Get directory Content and return a JSON-encoded string
func GetDirectoryContentInJSON(dir string) []byte {
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
		fullPath := []string{dir, "\\", file.Name()} // TODO: Check if it means anything to declare the slice before the for loop.
		c := Content{
			Name: name,
			Path: strings.Join(fullPath, ""),
			Size: file.Size(),
			Type: ext}

		listOfContent = append(listOfContent, c)
	}

	// Marshall a JSON-encoded string of the listOfContent
	mList, err := json.Marshal(listOfContent)
	if err != nil {
		log.Fatal(err)
	}

	return mList

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
