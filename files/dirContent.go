package dirContent

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

// Get directory Content and return
func getDirectoryContentInJSON(dir string) []byte {
	// Retrieve all the files in the input directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var listOfContent []Content

	// Create a Content struct for each file and convert it to a JSON object.
	for _, file := range files {
		ext := filepath.Ext(file.Name())
		name := strings.TrimSuffix(file.Name(), ext)
		c := Content{
			Name: name,
			Path: dir,
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
