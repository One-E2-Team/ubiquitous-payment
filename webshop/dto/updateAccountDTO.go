package dto

type UpdateAccountDTO struct {
	AccountID     string `json:"accountId"`
	Secret        string `json:"secret"`
	PaymentTypeId uint   `json:"paymentTypeId"`
}
