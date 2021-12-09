package main

import (
	"fmt"
	"ubiquitous-payment/psp-plugins/pspdto"
)

type plugin struct {
}

func (p plugin) Test() string {
	fmt.Println("Bit-in bit-out wasaaaaaaaaaa")
	return "bitups"
}

func (p plugin) SupportsPlanPayment() bool {
	return true
}

func (p plugin) ExecuteTransaction(data pspdto.TransactionDTO) (pspdto.TransactionCreatedDTO, error) {
	return pspdto.TransactionCreatedDTO{}, nil
}

var Plugin plugin
