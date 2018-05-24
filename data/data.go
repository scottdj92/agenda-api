package data

// Note a note struct
type Note struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags,omitempty"`
}

// Notes a basic note global var
var Notes []Note

func init() {
	Notes = append(Notes, Note{ID: "1", Title: "Walk the dog", Description: "Take Mojo out for a walk"})
	Notes = append(Notes, Note{ID: "2", Title: "Clean the house", Description: "Get some disinfectatant"})
}
