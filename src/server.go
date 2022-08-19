package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func getFunc() {
	return 1
}

func main() {
	fmt.Println("Hello World")

	router := mux.NewRouter()

	router.HandleFunc("/test", getFunc).Methods("GET")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":5000", handler))
}
