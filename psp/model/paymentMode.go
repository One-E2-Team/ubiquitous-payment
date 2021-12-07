package model

import "strings"

type PaymentMode int

const (
	ONE_TIME PaymentMode = iota
	RECURRING
)

func GetPaymentModeType(postType string) PaymentMode {
	if strings.ToLower(postType) == "RECURRING" {
		return RECURRING
	}
	return ONE_TIME
}
