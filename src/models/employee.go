package models

import "github.com/jinzhu/gorm"

type Employee struct {
	gorm.Model

	EmployeeId string `json:"EmployeeId"`

	Name string `json:"Name"`
}
