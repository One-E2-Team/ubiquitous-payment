package model

import "gorm.io/gorm"

type PaymentType struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}
