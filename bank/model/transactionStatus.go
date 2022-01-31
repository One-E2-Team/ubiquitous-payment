package model

type TransactionStatus string

const (
	PLACED    TransactionStatus = "PLACED"
	FULFILLED TransactionStatus = "FULFILLED"
	FAILED    TransactionStatus = "FAILED"
	ERROR     TransactionStatus = "ERROR"
)
