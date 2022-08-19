package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

type Usuario struct {
	gorm.Model

	GoogleId string `json:"GoogleId"`

	Nombre string `json:"Nombre"`

	Cargo string `json:"Cargo"`

	Correo string `json:"Correo"`
}

func GetFunc(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(&user)
}
