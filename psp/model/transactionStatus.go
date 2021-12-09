package model

type TransactionStatus int

const (
	WAITING TransactionStatus = iota
	FULLFILLED
	FAILED
	ERROR
)
