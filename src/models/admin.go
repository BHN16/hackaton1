package models

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model

	Name string `json:"Name"`

	Email string `json:"Email" gorm:"primaryKey"`

	Password string `json:"Password"`
}
