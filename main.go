package main

import (
	"log"
	"net/http"

	"./crud"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/note/{id}", crud.GetNote).Methods("GET")
	router.HandleFunc("/note/{id}", crud.CreateNote).Methods("POST")

	log.Println("server is running at port 1337")
	if err := http.ListenAndServe(":1337", router); err != nil {
		panic(err)
	}
}
