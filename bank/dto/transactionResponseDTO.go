package dto

import "time"

type TransactionResponseDTO struct {
	Amount                string    `json:"amount"`
	Currency              string    `json:"currency"`
	AcquirerAccountNumber string    `json:"acquirerAccountNumber"`
	IssuerPan             string    `json:"issuerPan"`
	Timestamp             time.Time `json:"timestamp"`
	TransactionStatus     string    `json:"transactionStatus"`
}
