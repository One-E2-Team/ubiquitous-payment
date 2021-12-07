package dto

type TransactionDTO struct {
	PspTransactionId   string
	OrderId            string
	PayeeId            string
	PayeeSecret        string
	Currency           string
	Amount             string
	ClientBusinessName string
	SuccessUrl         string
	FailUrl            string
	ErrorUrl           string
}

type TransactionCreatedDTO struct {
	TransactionId string
	RedirectUrl   string
}
