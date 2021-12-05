package model

type TransactionStatus int

const (
	FULLFILLED TransactionStatus = iota
	FAILED
	ERROR
)
