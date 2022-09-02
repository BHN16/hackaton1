package models

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model

	Name string `json:"Name"`

	Password string `json:"Password"`

	Role string `json:"Role"`
}
