package dto

type AccountResponseDTO struct {
	AccountNumber string                  `json:"accountNumber"`
	Amount        float32                 `json:"amount"`
	Secret        string                  `json:"secret"`
	CreditCards   []CreditCardResponseDTO `json:"creditCards"`
}

type CreditCardResponseDTO struct {
	Pan        string `json:"pan"`
	Cvc        string `json:"cvc"`
	HolderName string `json:"holderName"`
	ValidUntil string `json:"validUntil"`
}
