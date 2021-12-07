package model

type RecurringType int

const (
	MONTHLY RecurringType = iota
	YEARLY
)

func (rt RecurringType) String() string {
	return [...]string{"MONTHLY", "YEARLY"}[rt]
}
