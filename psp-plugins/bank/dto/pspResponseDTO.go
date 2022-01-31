package dto

type PspResponseDTO struct {
	PaymentId       string `json:"paymentId"`
	PaymentUrl      string `json:"paymentUrl"`
	PaymentCheckUrl string `json:"paymentCheckUrl"`
}
