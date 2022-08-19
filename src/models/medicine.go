package models

import "github.com/jinzhu/gorm"

type Mecine struct {
	gorm.Model

	MedicineId string `json:"MedicineId"`

	Name string `json:"Name"`
}
