package handlers

import (
	"encoding/json"
	"fmt"
	"hackaton/bd"
	"hackaton/models"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	var admin models.Admin
	var employee models.Employee
	var patient models.Patient

	// Models to know the table where the query will be executed and to store the result.
	/*
		bd.DB.Where("email = ?", user.Email).Find(&admin)
		bd.DB.Where("email = ?", user.Email).Find(&admin)
		bd.DB.Where("email = ?", user.Email).Find(&admin)
	*/
	bd.DB.Find(&admin, "Email = ?", user.Email)
	bd.DB.Find(&employee, "Email = ?", user.Email)
	bd.DB.Find(&patient, "Email = ?", user.Email)
	/*
		json.NewEncoder(w).Encode(&admin)
		json.NewEncoder(w).Encode(&employee)
		json.NewEncoder(w).Encode(&patient)*/

	// Check if the user in the database is admin, employee or patient and assign its role.
	if admin.Email != "" {
		if admin.Email == user.Email && comparePasswords(admin.Password, user.Password) {
			user.Role = "admin"
			//json.NewEncoder(w).Encode(map[string]bool{"response": true})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Username or Password"})
		}
	} else if employee.Email != "" {
		if employee.Email == user.Email && comparePasswords(employee.Password, user.Password) {
			user.Role = "employee"
			//json.NewEncoder(w).Encode(map[string]bool{"response": true})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Username or Password"})
		}
	} else if patient.Email != "" {
		if patient.Email == user.Email && comparePasswords(patient.Password, user.Password) {
			user.Role = "patient"
			//json.NewEncoder(w).Encode(map[string]bool{"response": true})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Username or Password"})
		}
	} else {
		json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Username or Password"})
	}

	// Generate the jwt with its role.

	validToken, err := GenerateJWT(user.Email, user.Role)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token models.Token
	token.Email = user.Email
	token.Role = user.Role
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)

	/*

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
	*/

}

func Register(w http.ResponseWriter, r *http.Request) {

	// email, user, password, tokenstring

	godotenv.Load(".env")


	var mySigningKey = []byte(os.Getenv("SECRET"))

	token, err := jwt.Parse(user.TokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	// Revisar los claims, estos tienen los atributos del token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Print(claims)
		if claims["Role"] == "admin" {
			fmt.Print("es admin")
		} else if claims["Role"] == "user" {
			fmt.Print("es usuario")
		}
	}


	var user models.UserAuthenticate

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}





	


	/*
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
	*/

}
