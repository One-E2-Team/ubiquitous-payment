package model

import (
	"gorm.io/gorm"
	"ubiquitous-payment/util"
)

type CreditCard struct {
	gorm.Model
	Pan        util.EncryptedString `json:"pan"`
	Cvc        util.EncryptedString `json:"cvc"`
	HolderName string               `json:"holderName" gorm:"not null"`
	ValidUntil string               `json:"validUntil" gorm:"not null"`
}
