package dto

import (
	"time"
	"ubiquitous-payment/bank/model"
)

type PaymentResponseDTO struct {
	MerchantOrderId   string                  `json:"merchantOrderId"`
	AcquirerOrderId   string                  `json:"acquirerOrderId"`
	AcquirerTimestamp time.Time               `json:"acquirerTimestamp"`
	PaymentId         string                  `json:"paymentId"`
	TransactionStatus model.TransactionStatus `json:"transactionStatus"`
}
