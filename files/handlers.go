package api

import (
	"encoding/json"
	"net/http"

	"github.com/paaff/Filly/error"
)

func BrowseHandler(w http.ResponseWriter, r *http.Request) *errorhandler.AppError {
	if r.Method == "POST" {

		decoder := json.NewDecoder(r.Body)
		var path Path
		jErr := decoder.Decode(&path)
		if jErr != nil {
			panic(jErr) // TODO: Correct error handling please.
		}

		// Browse from the POST form variable
		cont, err := GetDirectoryContentInJSON(path.Path)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cont)

	}
	return nil
}
