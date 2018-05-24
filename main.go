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

	log.Println("server is running at port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
