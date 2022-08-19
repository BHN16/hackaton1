package models

import "github.com/jinzhu/gorm"

type Patient struct {
	gorm.Model

	PatientId string `json:"PatientId"`

	Name string `json"Name"`
}
