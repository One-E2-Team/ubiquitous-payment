package dto

type TransactionStatusDTO int

const (
	WAITING TransactionStatusDTO = iota
	FULFILLED
	FAILED
	ERROR
)
