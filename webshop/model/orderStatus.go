package model

type OrderStatus int

const (
	PLACED OrderStatus = iota
	CANCELLED
	FULFILLED
	FAILED
	ERROR
)
