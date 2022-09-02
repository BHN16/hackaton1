package models

import "github.com/jinzhu/gorm"

type Patient struct {
	gorm.Model

	Name string `json:"Name"`

	Role string `json:"Role"`

}
