package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"hackaton/bd"
	"hackaton/models"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func checkRole(tokenString string) int {
	if tokenString == "" {
		return 0
	}

	godotenv.Load(".env")

	var mySigningKey = []byte(os.Getenv("SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return 3
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Print(claims)
		switch claims["Role"] {
		case "admin":
			return 1
		case "user":
			return 2
		default:
			return 0
		}
	}
	return 0
}

func setRole(user *models.User, admin *models.Admin, employee *models.Employee, patient *models.Patient) (string, error) {
	if admin.Email != "" {
		if user.Email == admin.Email && comparePasswords(admin.Password, user.Password) {
			return "admin", nil
		}
	} else if employee.Email != "" {
		if user.Email == employee.Email && comparePasswords(employee.Password, user.Password) {
			return "employee", nil
		}
	} else if patient.Email != "" {
		if user.Email == patient.Email && comparePasswords(patient.Password, user.Password) {
			return "patient", nil
		}
	}
	return "", errors.New("Invalid username or password")
}

func generateToken(user *models.User) (models.Token, error) {
	validToken, err := GenerateJWT(user.Email, user.Role)
	if err != nil {
		return models.Token{}, err
	}

	var token models.Token
	token.Email = user.Email
	token.Role = user.Role
	token.TokenString = validToken

	return token, nil
}

func InitializeDB(w http.ResponseWriter, r *http.Request) {
	bd.DB.AutoMigrate(&models.Admin{})
	bd.DB.AutoMigrate(&models.Employee{})
	bd.DB.AutoMigrate(&models.Medicine{})
	bd.DB.AutoMigrate(&models.Patient{})
	bd.DB.AutoMigrate(&models.Receipt{})
}

func Login(w http.ResponseWriter, r *http.Request) {

	var user models.User
	var admin models.Admin
	var employee models.Employee
	var patient models.Patient

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
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
	}

	token, err := generateToken(&user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(token)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {

<<<<<<< HEAD
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

	switch checkRole(user.TokenString) {
	case 1:
		fmt.Print("Es admin")
	case 2:
		fmt.Print("Es user")
	case 3:
		fmt.Print("Token expirado")
	default:
		fmt.Print("Error en el token")
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
