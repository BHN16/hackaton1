package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Name string `json:"Name"`

	Email string `json:"Email"`

	Password string `json:"Password"`

	Role string `json:"Role"`
}
