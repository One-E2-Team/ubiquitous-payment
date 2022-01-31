package main

import (
	"errors"
	"fmt"
	"ubiquitous-payment/psp-plugins/bank/transactions"
	"ubiquitous-payment/psp-plugins/pspdto"
)

type plugin struct {
	context map[string]string
}

func (p plugin) Test() string {
	fmt.Println("bankaaaaaaaaaaa")
	return "bank-ups"
}

func (p plugin) InitContextData(context map[string]string) {
	p.context = context
}

func (p plugin) SupportsPlanPayment() bool {
	return false
}

func (p plugin) ExecuteTransaction(data pspdto.TransactionDTO) (pspdto.TransactionCreatedDTO, error) {
	if data.PricingPlan {
		return pspdto.TransactionCreatedDTO{}, errors.New("bank does not support plan payment")
	}
	return transactions.PrepareTransaction(data, &p.context)
}

func (p plugin) CaptureTransaction(id string, plan bool) (bool, error) {
	if plan {
		return false, errors.New("bank does not allow for plan processing")
	}
	return transactions.CheckPaymentStatusSuccess(id)
}

var Plugin plugin

func main() {
	
}
