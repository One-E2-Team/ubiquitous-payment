package model

type PaymentMode int

const (
	ONE_TIME PaymentMode = iota
	RECURRING
)
