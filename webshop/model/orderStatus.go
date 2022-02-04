package model

type OrderStatus string

const (
	PLACED OrderStatus = "PLACED"
	CANCELLED OrderStatus = "CANCELLED"
	FULFILLED OrderStatus = "FULFILLED"
	FAILED OrderStatus = "FAILED"
	ERROR OrderStatus = "ERROR"
)
