package dto

type TransactionStatusDTO string

const (
	PLACED    TransactionStatusDTO = "PLACED"
	FULFILLED TransactionStatusDTO = "FULFILLED"
	FAILED    TransactionStatusDTO = "FAILED"
	ERROR     TransactionStatusDTO = "ERROR"
)
