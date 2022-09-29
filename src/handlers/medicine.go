package handlers

import (
	"encoding/json"
	"hackaton/bd"
	"hackaton/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMedicines(w http.ResponseWriter, r *http.Request) {

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

	var medicines []models.Medicine

	if role != 1 && role != 2 {
		json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Role"})
		return
	}

	bd.DB.Find(&medicines)

	json.NewEncoder(w).Encode(&medicines)

}

func GetMedicine(w http.ResponseWriter, r *http.Request) {

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

	var medicine models.Medicine

	bd.DB.First(&medicine, params["id"])

	json.NewEncoder(w).Encode(&medicine)

}

func PostMedicine(w http.ResponseWriter, r *http.Request) {

	var err_demo bool = true

	role, err := processCookie(r)

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			ErrorLogger.Println("No coockie")
			json.NewEncoder(w).Encode(map[string]string{"response": "No cookie"})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		ErrorLogger.Println("Bad request")
		json.NewEncoder(w).Encode(map[string]string{"response": "Bad request"})
		return
	}

	if role != 1 {
		ErrorLogger.Println("Invalid Role")
		json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Role"})
		return
	}

	var medicine models.Medicine
	err2 := json.NewDecoder(r.Body).Decode(&medicine)

	if err2 != nil {
		ErrorLogger.Println("Error en los datos recibidos", err2)
		http.Error(w, "Error en los datos recibidos"+err2.Error(), 400)
		return
	}

	if err_demo {
		w.WriteHeader(http.StatusBadRequest)
		dataLog, _ := json.Marshal(medicine)
		ErrorLogger.Println("Transaction Error", string(dataLog))
		json.NewEncoder(w).Encode(map[string]string{"response": "Transaction Error"})
		return
	}

	bd.DB.AutoMigrate(&models.Medicine{})

	bd.DB.Create(&medicine)

	json.NewEncoder(w).Encode(&medicine)

}
