package models

import "github.com/jinzhu/gorm"

type Receipt struct {
	gorm.Model

	PatientRefer int

	MedicineRefer int

	EmployeeRefer int

	patient Patient `gorm:"foreignKey:PatientRefer"`

	medicine Medicine `gorm:"foreignKey:MedicineRefer"`

	employee Employee `gorm:"foreignKey:EmployeeRefer"`
}
