package models

import "github.com/jinzhu/gorm"

type Medicine struct {
	gorm.Model

	Name string `json:"Name"`

	Cant uint `json:"Cant"`
}
type TempMedicine struct {

	Name string `json:"Name"`

	Cant uint `json:"Cant"`
}
