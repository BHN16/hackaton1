package handlers

import (
	"encoding/json"
	"hackaton/bd"
	"hackaton/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	var employees []models.Employee

	bd.DB.Find(&employees)

	json.NewEncoder(w).Encode(&employees)

}

func GetEmployee(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var employee models.Employee

	bd.DB.First(&employee, params["id"])

	json.NewEncoder(w).Encode(&employee)

}

func PostEmployee(w http.ResponseWriter, r *http.Request) {

	var employee models.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	bd.DB.AutoMigrate(&models.Employee{})

	bd.DB.Create(&employee)

	json.NewEncoder(w).Encode(&employee)

}
