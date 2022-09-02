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

	err2:= validateEntropy(admin.Password)

	if err2 != nil{
		http.Error(w, err2.Error(), 400)
		return
	}	
	
	admin.Password = hashAndSalt(admin.Password)

	
	bd.DB.AutoMigrate(&models.Admin{})

	bd.DB.Create(&admin)

	json.NewEncoder(w).Encode(&admin)

}

