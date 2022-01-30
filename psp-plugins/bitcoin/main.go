package main

import (
	"errors"
	"fmt"
	"ubiquitous-payment/psp-plugins/bitcoin/transactions"
	"ubiquitous-payment/psp-plugins/pspdto"
)

type plugin struct {
}

func (p plugin) Test() string {
	fmt.Println("Bit-in bit-out wasaaaaaaaaaa")
	return "bitups"
}

func (p plugin) InitContextData(_ map[string]string) {
	return
}

func (p plugin) SupportsPlanPayment() bool {
	return false
}

func (p plugin) ExecuteTransaction(data pspdto.TransactionDTO) (pspdto.TransactionCreatedDTO, error) {
	if data.PricingPlan {
		return pspdto.TransactionCreatedDTO{}, errors.New("bitcoin does not support plan payment")
	}
	return transactions.PrepareTransaction(data)
}

func (p plugin) CaptureTransaction(id string, plan bool) (bool, error) {
	if plan {
		return false, errors.New("bitcoin does not allow for plan processing")
	}
	return transactions.CaptureTransactionSuccess(id)
}

var Plugin plugin

func main() {

}
