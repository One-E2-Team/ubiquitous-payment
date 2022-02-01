package model

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	Amount            float32           `json:"amount"`
	AmountRsd         float32           `json:"amountRsd"`
	Currency          string            `json:"currency"`
	MerchantId        string            `json:"merchantId"`       //account number
	MerchantPassword  string            `json:"merchantPassword"` //account secret
	IssuerPan         string            `json:"issuerPan"`
	SuccessURL        string            `json:"successURL"`
	FailURL           string            `json:"failURL"`
	ErrorURL          string            `json:"errorURL"`
	MerchantOrderID   string            `json:"merchantOrderID"`
	MerchantTimestamp time.Time         `json:"merchantTimestamp"`
	PaymentId         string            `json:"paymentId"`
	PaymentUrlId      string            `json:"paymentUrlId"`
	TransactionStatus TransactionStatus `json:"transactionStatus"`
}

func (transaction *Transaction) GetURLByStatus() string {
	switch transaction.TransactionStatus {
	case FULFILLED:
		return transaction.SuccessURL
	case FAILED:
		return transaction.FailURL
	case ERROR:
		return transaction.ErrorURL
	default:
		return ""
	}
}
