package main

import (
	"hackaton/bd"
	"hackaton/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	godotenv.Load(".env")

	router := mux.NewRouter()

	bd.Connect()

	router.HandleFunc("/employees", handlers.GetEmployees).Methods("GET")

	router.HandleFunc("/employee/{id}", handlers.GetEmployee).Methods("GET")

	router.HandleFunc("/employee", handlers.PostEmployee).Methods("POST")

	router.HandleFunc("/medicines", handlers.GetMedicines).Methods("GET")

	router.HandleFunc("/medicine/{id}", handlers.GetMedicine).Methods("GET")

	router.HandleFunc("/medicine", handlers.PostMedicine).Methods("POST")

	router.HandleFunc("/patients", handlers.GetPatients).Methods("GET")

	router.HandleFunc("/patient/{id}", handlers.GetPatient).Methods("GET")

	router.HandleFunc("/patient", handlers.PostPatient).Methods("POST")

	router.HandleFunc("/admin", handlers.GetAdmin).Methods("GET")

	router.HandleFunc("/admin", handlers.PostAdmin).Methods("POST")

	router.HandleFunc("/login", handlers.Login).Methods("POST")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))

	defer bd.DB.Close()

}
