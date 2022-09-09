package models

import "github.com/jinzhu/gorm"

type Receipt struct {
	gorm.Model

	Cant uint `json:"Cant"`

	PatientRefer string `json:"Patient"`

	MedicineRefer uint `json:"Medicine"`

	EmployeeRefer string `json:"Employee"`

	patient Patient `gorm:"foreignKey:PatientRefer"`

	employee Employee `gorm:"foreignKey:EmployeeRefer"`

	medicines Medicine `gorm:"foreignKey:MedicinesRefer"`
}
