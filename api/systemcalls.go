package api

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/paaff/Filly/error"
)

// RootDir - Root directory global variable
var RootDir string

// Content - Directory content struct
type Content struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Size int64  `json:"size"`
	Type string `json:"type"`
}

// GetDirectoryContentInJSON - Get directory Content and return a JSON-encoded string
func GetDirectoryContentInJSON(dir string) ([]Content, *errorhandler.AppError) {

	fullDir := filepath.Join(RootDir, dir)
	// Retrieve all the files in the input directory
	files, err := ioutil.ReadDir(fullDir)
	if err != nil {
		return nil, &errorhandler.AppError{Error: err, Message: "Content cannot be found", Code: 404} // Status: Not Found
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

// DeleteSelectedContentFromDir - Delete user-chosen content from given directory.
func DeleteSelectedContentFromDir(removeDir string) *errorhandler.AppError {
	// Check if the path is valid.
	fullDir := filepath.Join(RootDir, removeDir)
	if _, err := os.Stat(fullDir); os.IsNotExist(err) {
		return &errorhandler.AppError{Error: err, Message: "Path does not exist", Code: 404} // Status: Not Found
	}

	// TODO: Revise the 401 error code.
	// Run through the paths and delete them.
	err := os.RemoveAll(removeDir)
	if err != nil {
		return &errorhandler.AppError{Error: err, Message: "Error in removing", Code: 401} // Status: Not Found
	}
	// No error means everything was deleted correctly.
	// TODO: Should make sure it deletes from valid paths within the system.

	return nil
}
