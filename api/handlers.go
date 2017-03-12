package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/paaff/Filly/error"
)

// BrowseHandler - Handler for browsing calls.
func BrowseHandler(w http.ResponseWriter, r *http.Request) *errorhandler.AppError {
	if r.Method == "POST" {
		path := apiCallDecode(r.Body)

		// Browse from the given path.
		cont, err := GetDirectoryContentInJSON(path.Source)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cont)

	}
	return nil
}

// RemoveHandler - Handler for removing folder/files.
func RemoveHandler(w http.ResponseWriter, r *http.Request) *errorhandler.AppError {
	if r.Method == "POST" {

		path := apiCallDecode(r.Body)

		// Remove from the given path.
		err := DeleteSelectedContentFromDir(path.Source)
		if err != nil {
			return err
		}
	}
	return nil
}

// Endpoint JSON decoder
func apiCallDecode(body io.ReadCloser) iocall {
	var path iocall
	decoder := json.NewDecoder(body)
	jErr := decoder.Decode(&path)
	if jErr != nil {
		panic(jErr) // TODO: Correct error handling please.
	}
	return path
}

// Path struct for unmarshalling JSON array of paths.
type iocall struct {
	Source      string `json:"src"`
	Destination string `json:"dst"`
}
