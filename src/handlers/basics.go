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
		fmt.Print(claims)
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

func generateToken(user *models.User) (string, error) {
	validToken, err := GenerateJWT(user.Email, user.Role)
	if err != nil {
		return "", err
	}

	/*

		var token models.Token
		token.Email = user.Email
		token.Role = user.Role
		token.TokenString = validToken
	*/
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

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   token,
			Expires: time.Now().Add(time.Minute * 5),
		})

		//w.Header().Set("Content-Type", "application/json")
		//json.NewEncoder(w).Encode(token)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {

	// email, user, password, tokenstring

	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
	}

	//var user models.UserAuthenticate

	//err := json.NewDecoder(r.Body).Decode(&user)

	godotenv.Load(".env")

	var mySigningKey = []byte(os.Getenv("SECRET"))

	token, err := jwt.Parse(c.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	// Revisar los claims, estos tienen los atributos del token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)

		switch checkRole(c.Value) {
		case 1:
			fmt.Println("Es admin")
		case 2:
			fmt.Println("Es doctor")
		case 3:
			fmt.Println("Es patient")
		case 4:
			fmt.Println("Token expirado")
		default:
			fmt.Println("Error en el token")
		}
	}

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	var user models.TemporalUser
	err2 := json.NewDecoder(r.Body).Decode(&user)
	if err2 != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	fmt.Println(user)

	switch checkRole(c.Value) {
	case 1:
		fmt.Println(user)

		switch user.Role {
		case "doctor":

			var doctor models.Employee
			doctor.Name = user.Name
			doctor.Codigo = user.Codigo
			doctor.Email = user.Email
			doctor.Password = user.Password
			//err3 := json.NewDecoder(r.Body).Decode(&doctor)

			err2 := validateEntropy(doctor.Password)

			if err2 != nil {
				http.Error(w, err2.Error(), 400)
				return
			}

			doctor.Password = hashAndSalt(doctor.Password)

			bd.DB.AutoMigrate(&models.Employee{})

			bd.DB.Create(&doctor)

			json.NewEncoder(w).Encode(&doctor)

		case "patient":

			var patient models.Patient

			patient.Name = user.Name
			patient.Email = user.Email
			patient.Password = user.Password

			err2 := validateEntropy(patient.Password)

			if err2 != nil {
				http.Error(w, err2.Error(), 400)
				return
			}

			patient.Password = hashAndSalt(patient.Password)

			patient = models.Patient(patient)

			bd.DB.AutoMigrate(&models.Patient{})

			bd.DB.Create(&patient)

			json.NewEncoder(w).Encode(&patient)

		default:
			json.NewEncoder(w).Encode(map[string]string{"response": "Invalid Role"})

		}

	case 2:
		fmt.Println("Es doctor")
	case 3:
		fmt.Println("Es patient")
	case 4:
		fmt.Println("Token expirado")
	default:
		fmt.Println("Error en el token")
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
