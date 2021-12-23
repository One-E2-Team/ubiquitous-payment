package dto

type CheckForPaymentDTO struct {
	PaymentCaptured bool   `json:"paymentCaptured"`
	SuccessUrl      string `json:"successUrl"`
}
