package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"golang.org/x/crypto/bcrypt"
	passwordvalidator "github.com/wagslane/go-password-validator"

)

const minEntropy = 60


func comparePasswords(hashedPwd string, plainPwd string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func validateEntropy(password string) error {
	const minEntropy = 60

	err2 := passwordvalidator.Validate(password, minEntropy)

	if err2 != nil {
		return err2
	}

	return nil

}

func hashAndSalt(pwd string) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
