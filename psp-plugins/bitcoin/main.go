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

func (p plugin) SupportsPlanPayment() bool {
	return false
}

func (p plugin) ExecuteTransaction(data pspdto.TransactionDTO) (pspdto.TransactionCreatedDTO, error) {
	if data.PricingPlan {
		return pspdto.TransactionCreatedDTO{}, errors.New("bitcoin does not support plan payment")
	}
	return transactions.PrepareTransaction(data)
}

var Plugin plugin

func main() {

}
