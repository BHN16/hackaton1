package handlers

import (
	"encoding/json"
	"hackaton/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var user models.TemporalUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error in the data"+err.Error(), 400)
		return
	}

	role, err2 := processCookie(r)

	if err2 != nil {
		if err2 == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"response": "No cookie"})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"response": "Bad request"})
		return
	}

	switch role {
	case 1:
		switch user.Role {
		case "doctor":

			toEncode, err := PostEmployee(&user)

			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			json.NewEncoder(w).Encode(toEncode)

		case "patient":

			toEncode, err := PostPatient(&user)

			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			json.NewEncoder(w).Encode(toEncode)

		default:
			json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Role"})
		}
	case 2:
		switch user.Role {
		case "patient":
			toEncode, err := PostPatient(&user)

			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			json.NewEncoder(w).Encode(toEncode)

		default:
			json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Role-Doctor"})
		}
	case 3:
		json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Role-Patient"})
	case 4:
		json.NewEncoder(w).Encode(map[string]string{"response": "Expired token"})
	default:
		json.NewEncoder(w).Encode(map[string]string{"response": "Token error"})
	}
}
