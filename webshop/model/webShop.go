package model

import "gorm.io/gorm"

type WebShop struct {
	gorm.Model
	Name 				string 				`json:"name" gorm:"unique;not null"`
	PSPAccessToken 		string 				`json:"pspAccessToken"`
	PaymentTypes 		[]PaymentType		`json:"paymentTypes" gorm:"many2many:webshop_paymenttypes;"`
}
