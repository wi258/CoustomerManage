package main

import (
	"crm/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/customer", routes.GetLists).Methods("GET")
	r.HandleFunc("/customer/{id}", routes.GetList).Methods("GET")
	r.HandleFunc("/customer", routes.CreateList).Methods("POST")
	r.HandleFunc("/customer/{id}", routes.UpdateLists).Methods("PUT")
	r.HandleFunc("/customer/{id}", routes.DeletLists).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", r))
}
