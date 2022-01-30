package model

import "gorm.io/gorm"

type CreditCard struct {
	gorm.Model
	Pan        string `json:"pan" gorm:"not null;unique"`
	Cvc        string `json:"cvc" gorm:"not null;unique"`
	HolderName string `json:"holderName" gorm:"not null"`
	ValidUntil string `json:"validUntil" gorm:"not null"`
}
