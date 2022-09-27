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

	router.HandleFunc("/", handlers.InitializeDB).Methods("POST")

	router.HandleFunc("/employees", handlers.GetEmployees).Methods("GET") //Admin

	router.HandleFunc("/employee/{id}", handlers.GetEmployee).Methods("GET") //Admin - (CurrentEmployee)?

	router.HandleFunc("/medicines", handlers.GetMedicines).Methods("GET") //Admin - Employees

	router.HandleFunc("/medicine/{id}", handlers.GetMedicine).Methods("GET") //Admin - Employees

	router.HandleFunc("/medicine", handlers.PostMedicine).Methods("POST") //Admin

	router.HandleFunc("/patients", handlers.GetPatients).Methods("GET") //Admin - Employees

	router.HandleFunc("/patient/{id}", handlers.GetPatient).Methods("GET") //Admin - Employees - (CurrentPatient)?

	router.HandleFunc("/receipts", handlers.GetReceipts).Methods("GET") //Admin - Employees - (CurrentPatient)?

	router.HandleFunc("/receipt", handlers.PostReceipt).Methods("POST") //Admin - Employees - (CurrentPatient)?

	router.HandleFunc("/login", handlers.Login).Methods("POST") //TO-DOS

	router.HandleFunc("/register", handlers.Register).Methods("POST") //ADMIN

	router.HandleFunc("/keepalive", handlers.Keepalive).Methods("GET")

	router.HandleFunc("/heartbeat", handlers.Heartbeat).Methods("GET")

	router.HandleFunc("/logtest", handlers.LogTest).Methods("GET")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))

	defer bd.DB.Close()

}
