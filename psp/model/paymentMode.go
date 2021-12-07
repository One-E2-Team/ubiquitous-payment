package model

import "strings"

type PaymentMode int

const (
	ONE_TIME PaymentMode = iota
	RECURRING
)

func GetPaymentMode(paymentMode string) PaymentMode {
	if strings.ToUpper(paymentMode) == "RECURRING" {
		return RECURRING
	}
	return ONE_TIME
}
