package dto

import "time"

type PspRequestDTO struct {
	MerchantId        string    `json:"merchantId"`
	MerchantPassword  string    `json:"merchantPassword"`
	Amount            float32   `json:"amount"`
	Currency          string    `json:"Currency"`
	MerchantOrderID   string    `json:"merchantOrderID"`
	MerchantTimestamp time.Time `json:"merchantTimestamp"`
	SuccessURL        string    `json:"successURL"`
	FailURL           string    `json:"failURL"`
	ErrorURL          string    `json:"errorURL"`
	Method            string    `json:"method"`
}
