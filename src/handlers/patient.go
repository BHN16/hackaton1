package handlers

import (
	"encoding/json"
	"hackaton/bd"
	"hackaton/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPatients(w http.ResponseWriter, r *http.Request) {

	var patients []models.Patient

	bd.DB.Find(&patients)

	json.NewEncoder(w).Encode(&patients)

}

func GetPatient(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var patient models.Patient

	bd.DB.First(&patient, params["id"])

	json.NewEncoder(w).Encode(&patient)

}

func PostPatient(w http.ResponseWriter, r *http.Request) {

	var patient models.Patient
	err := json.NewDecoder(r.Body).Decode(&patient)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	bd.DB.AutoMigrate(&models.Patient{})

	bd.DB.Create(&patient)

	json.NewEncoder(w).Encode(&patient)

}
