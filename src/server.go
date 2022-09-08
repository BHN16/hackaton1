package main

import (
	"hackaton/bd"
	"hackaton/handlers"
	"log"
	"net/http"

	//"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

/*
type person struct {
	Name     string
	LastName string
	Age      string
}*/

func main() {

	godotenv.Load(".env")

	router := mux.NewRouter()

	bd.Connect()

	/*
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			person := person{Name: "hola", LastName: "adios", Age: "12"}

			jsonr, _ := json.Marshal(person)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonr)

		}).Methods("GET")*/

	router.HandleFunc("/", handlers.InitializeDB).Methods("POST")

	router.HandleFunc("/employees", handlers.GetEmployees).Methods("GET") //Admin

	router.HandleFunc("/employee/{id}", handlers.GetEmployee).Methods("GET") //Admin - CurrentEmployee

	//router.HandleFunc("/employee", handlers.PostEmployee).Methods("POST")   //

	router.HandleFunc("/medicines", handlers.GetMedicines).Methods("GET") //Admin - Employees

	router.HandleFunc("/medicine/{id}", handlers.GetMedicine).Methods("GET") //Admin - Employees

	router.HandleFunc("/medicine", handlers.PostMedicine).Methods("POST") //Admin

	router.HandleFunc("/patients", handlers.GetPatients).Methods("GET") //Admin - Employees

	router.HandleFunc("/patient/{id}", handlers.GetPatient).Methods("GET") //Admin - Employees - CurrentPatient

	//ruta para crear recetas --> Admin - Employees

	//ruta para ver todas las recetas de un usuario --> Admin - Employee - CurrentPatient

	//router.HandleFunc("/patient", handlers.PostPatient).Methods("POST")

	router.HandleFunc("/admin", handlers.GetAdmin).Methods("GET") //DEBEMOS BORRARLO

	router.HandleFunc("/admin", handlers.PostAdmin).Methods("POST") //DEBEMOS BORRARLO

	router.HandleFunc("/login", handlers.Login).Methods("POST") //TO-DOS

	router.HandleFunc("/register", handlers.Register).Methods("POST") //ADMIN

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))

	defer bd.DB.Close()

}
