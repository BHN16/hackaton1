package handlers

import (
	"encoding/json"
	"hackaton/bd"
	"hackaton/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPatients(w http.ResponseWriter, r *http.Request) {

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

	if role != 1 && role != 2 {
		json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Role"})
		return
	}

	var patients []models.Patient

	bd.DB.Find(&patients)

	json.NewEncoder(w).Encode(&patients)

}

func GetPatient(w http.ResponseWriter, r *http.Request) {

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

	if role != 1 && role != 2 {
		json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Role"})
		return
	}

	params := mux.Vars(r)

	var patient models.Patient

	bd.DB.First(&patient, params["id"])

	json.NewEncoder(w).Encode(&patient)

}

func PostPatient(user *models.TemporalUser) (*models.Patient, error) {

	var patient models.Patient

	patient.Name = user.Name
	patient.Email = user.Email
	patient.Password = user.Password

	err := validateEntropy(patient.Password)

	if err != nil {
		return nil, err
	}

	patient.Password = hashAndSalt(patient.Password)

	// PARTE REDUNDANTE, NO SÉ QUE ES
	// patient = models.Patient(patient)

	bd.DB.AutoMigrate(&models.Patient{})

	bd.DB.Create(&patient)

	return &patient, nil
	/*err := json.NewDecoder(r.Body).Decode(&patient)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	err2 := validateEntropy(patient.Password)

	if err2 != nil {
		http.Error(w, err2.Error(), 400)
		return
	}

	patient.Password = hashAndSalt(patient.Password)

	bd.DB.AutoMigrate(&models.Patient{})

	bd.DB.Create(&patient)

	json.NewEncoder(w).Encode(&patient)*/

}
