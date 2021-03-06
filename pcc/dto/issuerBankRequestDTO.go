package dto

import (
	"time"
	"ubiquitous-payment/pcc/model"
)

type IssuerBankRequestDTO struct {
	AcquirerTransactionId uint              `json:"acquirerTransactionId"`
	AcquirerTimestamp     time.Time         `json:"acquirerTimestamp"`
	AcquirerPanPrefix     string            `json:"acquirerPanPrefix"`
	MerchantId            string            `json:"merchantId"`
	Amount                float32           `json:"amount"`
	Currency              string            `json:"currency"`
	IssuerPAN             string            `json:"issuerPan"`
	IssuerCVC             string            `json:"issuerCvc"`
	IssuerValidUntil      string            `json:"issuerValidUntil"`
	IssuerHolderName      string            `json:"issuerHolderName"`
	IssuerOrderId         uint              `json:"issuerOrderId"`
	IssuerTimestamp       time.Time         `json:"issuerTimestamp"`
	OrderStatus           model.OrderStatus `json:"orderStatus"`
}
