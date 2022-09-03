package handlers

import (
	"encoding/json"
	"hackaton/bd"
	"hackaton/models"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	switch user.Role {
	case "Admin":
		var admin models.Admin
		bd.DB.First(&admin)

		if admin.Email == user.Email && comparePasswords(admin.Password, user.Password) {
			json.NewEncoder(w).Encode(map[string]bool{"response": true})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Username or Password"})
		}

	case "Employee":
		var employee models.Employee
		bd.DB.Find(&employee, "Email = ?", user.Email)

		if employee.Email == user.Email && comparePasswords(employee.Password, user.Password) {
			json.NewEncoder(w).Encode(map[string]bool{"response": true})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Username or Password"})
		}

	case "Patient":
		var patient models.Patient
		bd.DB.Find(&patient, "Email = ?", user.Email)

		if patient.Email == user.Email && comparePasswords(patient.Password, user.Password) {
			json.NewEncoder(w).Encode(map[string]bool{"response": true})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Username or Password"})
		}

	default:
		json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Role"})

	}

}

func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	switch user.Role {
	case "Employee":
		var employee models.Employee

		err2 := validateEntropy(user.Password)

		if err2 != nil {
			http.Error(w, err2.Error(), 400)
			return
		}

		user.Password = hashAndSalt(user.Password)

		employee = models.Employee(user)

		bd.DB.AutoMigrate(&models.Employee{})

		bd.DB.Create(&employee)

		json.NewEncoder(w).Encode(&employee)

	case "Patient":
		err2 := validateEntropy(user.Password)

		if err2 != nil {
			http.Error(w, err2.Error(), 400)
			return
		}

		user.Password = hashAndSalt(user.Password)

		var patient models.Patient

		patient = models.Patient(user)

		bd.DB.AutoMigrate(&models.Patient{})

		bd.DB.Create(&patient)

		json.NewEncoder(w).Encode(&patient)
	default:
		json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Role"})

	}

}
