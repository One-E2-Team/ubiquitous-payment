package dto

import "time"

type MyOrderDTO struct {
	OrderId    		uint 		`json:"orderId"`
	Timestamp 		time.Time	`json:"timestamp"`
	ProductName		string		`json:"productName"`
	ProductPrice	float32		`json:"productPrice"`
	Currency		string		`json:"currency"`
	PaymentType		string 		`json:"paymentType"`
	PSPId			string		`json:"pspId"`
	OrderStatus 	string		`json:"orderStatus"`
	NumberOfInstallments uint	`json:"numberOfInstallments"`
	DelayedInstallments uint 	`json:"delayedInstallments"`
	RecurringType 	string		`json:"recurringType"`
}
