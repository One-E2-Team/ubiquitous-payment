package model

import (
	"gorm.io/gorm"
	"time"
	"ubiquitous-payment/util"
)

type PccOrder struct {
	gorm.Model
	AcquirerTransactionId uint                 `json:"acquirerTransactionId"`
	AcquirerTimestamp     time.Time            `json:"acquirerTimestamp"`
	AcquirerPanPrefix     string               `json:"acquirerPanPrefix"`
	MerchantId            string               `json:"merchantId"`
	Amount                float32              `json:"amount"`
	Currency              string               `json:"currency"`
	IssuerPAN             util.EncryptedString `json:"issuerPan"`
	IssuerCVC             util.EncryptedString `json:"issuerCvc"`
	IssuerValidUntil      string               `json:"issuerValidUntil"`
	IssuerHolderName      string               `json:"issuerHolderName"`
	IssuerOrderId         uint                 `json:"issuerOrderId"`
	IssuerTimestamp       time.Time            `json:"issuerTimestamp"`
	OrderStatus           OrderStatus          `json:"orderStatus"`
}
