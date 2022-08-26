package main

import (
	"hackaton/bd"
	"hackaton/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()

	bd.Connect()

	router.HandleFunc("/employees", handlers.GetEmployees).Methods("GET")

	router.HandleFunc("/employee/{id}", handlers.GetEmployee).Methods("GET")

	router.HandleFunc("/employee", handlers.PostEmployee).Methods("POST")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":5000", handler))

	defer bd.DB.Close()

}
