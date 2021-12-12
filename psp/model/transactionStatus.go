package model

type TransactionStatus int

const (
	WAITING TransactionStatus = iota
	FULFILLED
	FAILED
	ERROR
)

func (status *TransactionStatus) ToString() string {
	switch *status {
	case FULFILLED:
		return "success"
	case FAILED:
		return "fail"
	default:
		return "error"
	}
}
