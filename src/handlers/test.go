package handlers

import (
	"encoding/json"
	"net/http"
)

func GetFunc(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(&user)
}
