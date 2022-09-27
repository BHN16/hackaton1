package handlers

import (
	"errors"
	"fmt"
	"hackaton/models"
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
