package dto

import (
	"time"
)

type PaymentResponseDTO struct {
	MerchantOrderId   string               `json:"merchantOrderId"`
	AcquirerOrderId   string               `json:"acquirerOrderId"`
	AcquirerTimestamp time.Time            `json:"acquirerTimestamp"`
	PaymentId         string               `json:"paymentId"`
	TransactionStatus TransactionStatusDTO `json:"transactionStatus"`
}
