package handlers

import (
	"encoding/json"
	"hackaton/bd"
	"hackaton/models"
	"net/http"
)

func GetReceipt(w http.ResponseWriter, r *http.Request) {

	var receipts []models.Receipt

	bd.DB.Find(&receipts)

	json.NewEncoder(w).Encode(&receipts)

}

func PostReceipt(w http.ResponseWriter, r *http.Request) {

}
