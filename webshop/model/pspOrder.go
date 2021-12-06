package model

import (
	"gorm.io/gorm"
	"time"
)

type PSPOrder struct {
	gorm.Model
	PSPId			string		   `json:"pspId" gorm:"not null;unique"`
	Timestamp   	time.Time      `json:"timestamp"`
	OrderId			uint           `json:"orderId" gorm:"not null"`
	OrderStatus		OrderStatus	   `json:"orderStatus"`
}
