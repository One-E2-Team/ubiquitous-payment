package model

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	Amount            float32           `json:"amount"`
	Currency          string            `json:"Currency"`
	MerchantId        string            `json:"merchantId"`       //account number
	MerchantPassword  string            `json:"merchantPassword"` //account secret
	SuccessURL        string            `json:"successURL"`
	FailURL           string            `json:"failURL"`
	ErrorURL          string            `json:"errorURL"`
	MerchantOrderID   string            `json:"merchantOrderID"`
	MerchantTimestamp time.Time         `json:"merchantTimestamp"`
	PaymentId         string            `json:"paymentId"`
	PaymentUrlId      string            `json:"paymentUrlId"`
	TransactionStatus TransactionStatus `json:"transactionStatus"`
}
