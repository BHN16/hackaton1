package handlers

import (
	"encoding/json"
	"hackaton/bd"
	"hackaton/models"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var cur_admin models.Admin
	var admin models.Admin
	err := json.NewDecoder(r.Body).Decode(&cur_admin)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	bd.DB.First(&admin)

	if admin.Name == cur_admin.Name && comparePasswords(admin.Password, cur_admin.Password) {
		json.NewEncoder(w).Encode(map[string]bool{"response": true})
	} else {
		json.NewEncoder(w).Encode(map[string]bool{"response": false})
	}

}
