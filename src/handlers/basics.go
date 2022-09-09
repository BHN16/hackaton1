package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"hackaton/bd"
	"hackaton/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func getEmail(tokenString string) interface{} {
	if tokenString == "" {
		return ""
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
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["Email"]
	}
	return ""
}

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
		return 4
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		switch claims["Role"] {
		case "admin":
			return 1
		case "doctor":
			return 2
		case "patient":
			return 3
		default:
			return 0
		}
	}
	return 0
}

func processCookie(r *http.Request) (int, error) {
	c, err := r.Cookie("token")
	if err != nil {
		return -1, err
	} else {
		return checkRole(c.Value), nil
	}
}

func setRole(user *models.User, admin *models.Admin, employee *models.Employee, patient *models.Patient) (string, error) {
	if admin.Email != "" {
		if user.Email == admin.Email && comparePasswords(admin.Password, user.Password) {
			return "admin", nil
		}
	} else if employee.Email != "" {
		if user.Email == employee.Email && comparePasswords(employee.Password, user.Password) {
			return "doctor", nil
		}
	} else if patient.Email != "" {
		if user.Email == patient.Email && comparePasswords(patient.Password, user.Password) {
			return "patient", nil
		}
	}
	return "", errors.New("Invalid username or password")
}

func generateToken(user *models.User) (string, error) {
	validToken, err := GenerateJWT(user.Email, user.Role)
	if err != nil {
		return "", err
	}
	return validToken, nil
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
