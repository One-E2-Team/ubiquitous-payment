package model

import (
	"gorm.io/gorm"
	"time"
)

type PccOrder struct {
	gorm.Model
	AcquirerTransactionId uint      `json:"acquirerTransactionId"`
	AcquirerTimestamp     time.Time `json:"acquirerTimestamp"`
	Amount                float32   `json:"amount"`
	Currency              string    `json:"currency"`
	IssuerPan             string    `json:"issuerPan"`
	IssuerCvc             string    `json:"issuerCvc"`
	IssuerValidUntil      string    `json:"issuerValidUntil"`
	IssuerHolderName      string    `json:"issuerHolderName"`
	IssuerTransactionId   uint      `json:"issuerTransactionId"`
	IssuerTimestamp       time.Time `json:"issuerTimestamp"`
}
