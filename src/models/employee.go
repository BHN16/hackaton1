package models

import "github.com/jinzhu/gorm"

type Employee struct {
	gorm.Model

	Name string `json:"Name"`
}
