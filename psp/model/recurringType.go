package model

import "strings"

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

func GetRecurringString(recurringType RecurringType) string {
	if recurringType == MONTHLY{
		return "MONTHLY"
	}
	return "YEARLY"
}
