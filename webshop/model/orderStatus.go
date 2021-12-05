package model

type OrderStatus int

const (
	PLACES OrderStatus = iota
	CANCELLED
	FULFILLED
	FAILED
	ERROR
)
