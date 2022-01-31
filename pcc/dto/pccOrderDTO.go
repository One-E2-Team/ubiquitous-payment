package dto

import "time"

type PccOrderDto struct {
	AcquirerTransactionId uint      `json:"acquirerTransactionId"`
	AcquirerTimestamp     time.Time `json:"acquirerTimestamp"`
	AcquirerPanPrefix     string    `json:"acquirerPanPrefix"`
	MerchantId            string    `json:"merchantId"`
	Amount                float32   `json:"amount"`
	Currency              string    `json:"currency"`
	IssuerPAN             string    `json:"issuerPan"`
	IssuerCVC             string    `json:"issuerCvc"`
	IssuerValidUntil      string    `json:"issuerValidUntil"`
	IssuerHolderName      string    `json:"issuerHolderName"`
}
