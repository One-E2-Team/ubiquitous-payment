package model

type TransactionStatus int

const (
	WAITING TransactionStatus = iota
	FULFILLED
	FAILED
	ERROR
)
