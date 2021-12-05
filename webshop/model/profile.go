package model

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	FirstName 		string 		`json:"firstName" gorm:"not null"`
	LastName 		string 		`json:"lastName" gorm:"not null"`
	Email 			string 		`json:"email" gorm:"unique;not null"`
}
