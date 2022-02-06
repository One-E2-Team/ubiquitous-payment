package model

import (
	"gorm.io/gorm"
	"ubiquitous-payment/util"
)

type Account struct {
	gorm.Model
	AccountID     util.EncryptedString `json:"accountId"`
	Secret        util.EncryptedString `json:"secret"`
	PaymentTypeId uint                 `json:"paymentTypeId" gorm:"not null"`
	ProfileId     uint                 `json:"profileId" gorm:"not null"`
}
