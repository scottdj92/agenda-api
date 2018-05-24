package crud

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"../data"
)

// CreateNote creates a single note in our db
func CreateNote(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var note data.Note
	_ = json.NewDecoder(req.Body).Decode(&note)

	note.ID = params["id"]
	// append to DB here
	json.NewEncoder(res).Encode(note)
}
