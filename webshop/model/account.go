package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountID     string `json:"accountId" gorm:"not null"`
	Secret        string `json:"secret"`
	PaymentTypeId uint   `json:"paymentTypeId" gorm:"not null"`
	ProfileId     uint   `json:"profileId" gorm:"not null"`
}
