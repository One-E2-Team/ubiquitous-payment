package main

import (
	"fmt"
	"ubiquitous-payment/psp-plugins/paypal/transactions"
	"ubiquitous-payment/psp-plugins/pspdto"
)

type plugin struct {
}

func (p plugin) Test() string {
	fmt.Println("Plug-in plug-out wasaaaaaaaaaa")
	return "ups"
}

func (p plugin) SupportsPlanPayment() bool {
	return true
}

func (p plugin) ExecuteTransaction(data pspdto.TransactionDTO) (pspdto.TransactionCreatedDTO, error) {
	if data.PricingPlan {
		return transactions.ExecuteSubscription(data)
	} else {
		return transactions.ExecuteOrder(data)
	}
}

var Plugin plugin
