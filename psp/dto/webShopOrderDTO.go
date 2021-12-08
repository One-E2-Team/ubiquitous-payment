package dto

import "time"

type WebShopOrderDTO struct {
	PspOrderId        string                 `json:"pspOrderId"`
	Amount            float32                `json:"amount"`
	Currency          string                 `json:"currency"`
	PaymentMode       string                 `json:"paymentMode,omitempty"`
	IsSubscription    bool                   `json:"isSubscription"`
	RecurringType     string                 `json:"recurringType,omitempty"`
	RecurringTimes    string                 `json:"recurringTimes,omitempty"`
	PaymentTo         map[string][]string    `json:"paymentTo,omitempty"`
	SuccessUrl        string                 `json:"successUrl"`
	FailUrl           string                 `json:"failUrl"`
	ErrorUrl          string                 `json:"errorUrl"`
	MerchantTimestamp time.Time              `json:"merchantTimestamp"`
	MerchantOrderId   string                 `json:"merchantOrderId"`
}
