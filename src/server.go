package main

import (
	"hackaton/bd"
	"hackaton/handlers"
	"log"
	"net/http"
	"os"

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

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), handler))

	defer bd.DB.Close()

}
