package model

type OrderStatus string

const (
	PLACED OrderStatus = "PLACED"
	FULFILLED  OrderStatus = "FULFILLED"
	FAILED OrderStatus = "FAILED"
	ERROR OrderStatus = "ERROR"
)
