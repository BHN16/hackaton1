package handlers

import (
	"encoding/json"
	"fmt"
	"hackaton/bd"
	"hackaton/models"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var user models.User
	var admin models.Admin
	var employee models.Employee
	var patient models.Patient

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error in the data"+err.Error(), 400)
		return
	}

	bd.DB.Find(&admin, "Email = ?", user.Email)
	bd.DB.Find(&employee, "Email = ?", user.Email)
	bd.DB.Find(&patient, "Email = ?", user.Email)

	assignedRole, err := setRole(&user, &admin, &employee, &patient)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Username or Password"})
	} else {
		user.Role = assignedRole
		fmt.Println("Rol asignado al momento de hacer login: ", user.Role)
	}

	fmt.Println(user)

	token, err := generateToken(&user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	} else {

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   token,
			Expires: time.Now().Add(time.Minute * 5),
		})
	}
}
