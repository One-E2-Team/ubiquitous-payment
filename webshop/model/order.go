package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Timestamp 		 	time.Time      `json:"timestamp"`
	BuyerProfileId   	uint           `json:"buyerProfileId" gorm:"not null"`
	ProductId		   	uint           `json:"productId" gorm:"not null"`
	PaymentTypeId		uint		   `json:"paymentTypeId" gorm:"not null"`
}