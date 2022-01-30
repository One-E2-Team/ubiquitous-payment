package model

import "gorm.io/gorm"

type ClientAccount struct {
	gorm.Model
	AccountNumber string       `json:"accountNumber" gorm:"not null;unique"`
	Amount        float32      `json:"amount" gorm:"not null"`
	Secret        string       `json:"secret" gorm:"not null;unique"`
	IsActive      bool         `json:"isActive" gorm:"not null"`
	CreditCards   []CreditCard `json:"creditCards" gorm:"many2many:account_cards;"`
}
