package handlers

import (
	"encoding/json"
	"hackaton/bd"
	"hackaton/models"
	"net/http"
)

func GetAdmin(w http.ResponseWriter, r *http.Request) {

	var admin []models.Admin

	bd.DB.First(&admin)

	json.NewEncoder(w).Encode(&admin)

}

func PostAdmin(w http.ResponseWriter, r *http.Request) {

	var admin models.Admin
	err := json.NewDecoder(r.Body).Decode(&admin)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	bd.DB.AutoMigrate(&models.Admin{})

	bd.DB.Create(&admin)

	json.NewEncoder(w).Encode(&admin)

}
