package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

const minEntropy = 60

func comparePasswords(hashedPwd string, plainPwd string) bool {
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
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
