package dto

type PaymentDetailsDTO struct {
	Amount    float32 `json:"amount"`
	AmountRsd float32 `json:"amountRsd"`
	Currency  string  `json:"currency"`
}
