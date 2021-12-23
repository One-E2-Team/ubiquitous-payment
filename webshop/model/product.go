package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name                string        `json:"name" gorm:"not null"`
	Price               float32       `json:"price"`
	Currency            string        `json:"currency" gorm:"not null"`
	Description         string        `json:"description"`
	MediaPath           string        `json:"mediaPath"`
	IsActive            bool          `json:"isActive" gorm:"not null"`
	NumOfInstallments   uint          `json:"numOfInstallments"`
	DelayedInstallments uint          `json:"delayedInstallments"`
	MerchantProfileId   uint          `json:"merchantProfileId"`
	RecurringType       RecurringType `json:"recurringType"`
}

func (product *Product) Deactivate() {
	product.IsActive = false
}

func (product *Product) Update(updatedProduct Product) {
	product.Name = updatedProduct.Name
	product.Currency = updatedProduct.Currency
	product.Description = updatedProduct.Description
	product.NumOfInstallments = updatedProduct.NumOfInstallments
}
