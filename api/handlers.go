package api

import (
	"encoding/json"
	"net/http"

	"github.com/paaff/Filly/error"
)

// BrowseHandler - Handler for browsing calls.
func BrowseHandler(w http.ResponseWriter, r *http.Request) *errorhandler.AppError {
	if r.Method == "POST" {

		decoder := json.NewDecoder(r.Body)
		var path io
		jErr := decoder.Decode(&path)
		if jErr != nil {
			panic(jErr) // TODO: Correct error handling please.
		}

		// Browse from the POST form variable
		cont, err := GetDirectoryContentInJSON(path.Source)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cont)

	}
	return nil
}

// Path struct for unmarshalling JSON array of paths.
type io struct {
	Source      string `json:"src"`
	Destination string `json:"dst"`
}
