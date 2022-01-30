package transactions

import (
	"errors"
	"ubiquitous-payment/psp-plugins/pspdto"
)

func PrepareTransaction(data pspdto.TransactionDTO, context *map[string]string) (pspdto.TransactionCreatedDTO, error) {
	return pspdto.TransactionCreatedDTO{}, errors.New("unimplemented")
}

func CheckPaymentStatusSuccess(id string) (bool, error) {
	return false, errors.New("unimplemented")
}
