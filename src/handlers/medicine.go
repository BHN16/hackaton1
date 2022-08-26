package handlers

import (
	"encoding/json"
	"hackaton/bd"
	"hackaton/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMedicines(w http.ResponseWriter, r *http.Request) {

	var medicines []models.Medicine

	bd.DB.Find(&medicines)

	json.NewEncoder(w).Encode(&medicines)

}

func GetMedicine(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var medicine models.Medicine

	bd.DB.First(&medicine, params["id"])

	json.NewEncoder(w).Encode(&medicine)

}

func PostMedicine(w http.ResponseWriter, r *http.Request) {

	var medicine models.Medicine
	err := json.NewDecoder(r.Body).Decode(&medicine)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	bd.DB.AutoMigrate(&models.Medicine{})

	bd.DB.Create(&medicine)

	json.NewEncoder(w).Encode(&medicine)

}
