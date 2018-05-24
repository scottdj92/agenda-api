package crud

import (
	"encoding/json"
	"net/http"

	"../data"
	"github.com/gorilla/mux"
)

// GetNote gets a single note from our data structure
func GetNote(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range data.Notes {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
		}
	}
}
