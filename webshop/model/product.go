package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name 					string 				`json:"name" gorm:"not null"`
	Price 					float32 			`json:"price"`
	Currency 				string 				`json:"currency" gorm:"not null"`
	Description 			string 				`json:"description"`
	MediaPath 				string 				`json:"mediaPath"`
	IsActive 				bool	 			`json:"isActive" gorm:"not null"`
	NumOfInstallments 		uint 				`json:"numOfInstallments"`
	MerchantProfileId		uint				`json:"merchantProfileId"`
}
