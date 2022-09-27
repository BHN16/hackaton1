package handlers

import (
	"encoding/json"
	"fmt"
	"hackaton/bd"
	"hackaton/models"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func InitializeDB(w http.ResponseWriter, r *http.Request) {
	bd.DB.AutoMigrate(&models.Admin{})
	bd.DB.AutoMigrate(&models.Employee{})
	bd.DB.AutoMigrate(&models.Medicine{})
	bd.DB.AutoMigrate(&models.Patient{})
	bd.DB.AutoMigrate(&models.Receipt{})
}

func Keepalive(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"INFO": "ALIVE"})
}

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	result := bd.DB.First(&employee)
	if result.Error != nil {
		w.WriteHeader(500)
		log.Println("ERROR: NO CONECTION DATABASE") // LOG level: (Warning, Error, Info, Debug)+..
		json.NewEncoder(w).Encode(map[string]string{"ERROR": "NO CONECTION DATABASE"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"INFO": "DATABASE LIVE"})
}

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

func processCookie(r *http.Request) (int, error) {
	c, err := r.Cookie("token")
	if err != nil {
		return -1, err
	} else {
		return checkRole(c.Value), nil
	}
}

func generateToken(user *models.User) (string, error) {
	validToken, err := GenerateJWT(user.Email, user.Role)
	if err != nil {
		return "", err
	}
	return validToken, nil
}

func LogTest(w http.ResponseWriter, r *http.Request) {
	log.Printf("ERROR test")
}
