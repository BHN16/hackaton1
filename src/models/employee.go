package models

import "github.com/jinzhu/gorm"

type Employee struct {
	gorm.Model

	Name string `json:"Name"`

	Email string `json:"Email"`

	Password string `json:"Password"`

	Role string `json:"Role"`
}
