package handlers

import (
	"encoding/json"
	"hackaton/bd"
	"hackaton/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	role, err := processCookie(r)

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"response": "No cookie"})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"response": "Bad request"})
		return
	}

	if role != 1 {
		json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Role"})
		return
	}

	var employees []models.Employee

	bd.DB.Find(&employees)

	json.NewEncoder(w).Encode(&employees)

}

func GetEmployee(w http.ResponseWriter, r *http.Request) {

	role, err := processCookie(r)

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"response": "No cookie"})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"response": "Bad request"})
		return
	}
	// CHECK FOR THE CASE OF CURRENTEMPLOYEE
	if role != 1 {
		json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Role"})
		return
	}

	params := mux.Vars(r)

	var employee models.Employee

	bd.DB.First(&employee, params["id"])

	json.NewEncoder(w).Encode(&employee)

}

func PostEmployee(user *models.TemporalUser) (*models.Employee, error) {

	var doctor models.Employee

	doctor.Name = user.Name
	doctor.Codigo = user.Codigo
	doctor.Email = user.Email
	doctor.Password = user.Password

	err := validateEntropy(doctor.Password)

	if err != nil {
		return nil, err
	}

	doctor.Password = hashAndSalt(doctor.Password)

	bd.DB.AutoMigrate(&models.Employee{})

	bd.DB.Create(&doctor)

	return &doctor, nil
	// json.NewEncoder(w).Encode(&employee)
}
