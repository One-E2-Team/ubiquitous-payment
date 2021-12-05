package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Transaction struct {
	ID                    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	PSPId                 string             `json:"pspId"`
	WebShopID             string             `json:"webShopID"`
	Amount                float32            `json:"amount"`
	Currency              string             `json:"Currency"`
	SuccessURL            string             `json:"successURL"`
	FailURL               string             `json:"failURL"`
	ErrorURL              string             `json:"errorURL"`
	MerchantOrderID       int                `json:"merchantOrderID"`
	MerchantTimestamp     time.Time          `json:"merchantTimestamp"`
	PaymentMode           PaymentMode        `json:"paymentMode"`
	Recurring             *Recurring         `json:"recurring"`
	TransactionStatus     TransactionStatus  `json:"transactionStatus"`
	AvailablePaymentTypes []PaymentType      `json:"availablePaymentTypes"`
	SelectedPaymentType   PaymentType        `json:"selectedPaymentType"`
	MerchantAccounts      []Account          `json:"merchant_accounts"`
}
