package model

import (
	"gorm.io/gorm"
	"ubiquitous-payment/util"
)

type ClientAccount struct {
	gorm.Model
	AccountNumber util.EncryptedString `json:"accountNumber"`
	Amount        float32              `json:"amount" gorm:"not null"`
	Secret        util.EncryptedString `json:"secret"`
	IsActive      bool                 `json:"isActive" gorm:"not null"`
	CreditCards   []CreditCard         `json:"creditCards" gorm:"many2many:account_cards;"`
}
