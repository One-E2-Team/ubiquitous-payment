package model

import (
	"strings"
	"ubiquitous-payment/psp-plugins/pspdto"
)

type RecurringType int

const (
	MONTHLY RecurringType = iota
	YEARLY
)

func GetRecurringType(recurringType string) RecurringType {
	if strings.ToUpper(recurringType) == "YEARLY" {
		return YEARLY
	}
	return MONTHLY
}

func GetInstallmentUnitByRecurringType(recurringType RecurringType) pspdto.InstallmentUnit {
	switch recurringType {
	case YEARLY:
		return pspdto.Year
	default:
		return pspdto.Month
	}
}
