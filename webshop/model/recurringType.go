package model

type RecurringType string

const (
	MONTHLY RecurringType = "MONTHLY"
	YEARLY  RecurringType = "YEARLY"
)

func GetRecurringType(str string) RecurringType {
	switch str {
	case "MONTHLY":
		return MONTHLY
	case "YEARLY":
		return YEARLY
	default:
		return MONTHLY
	}
}
