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

func (plugin) Name() string {
	return "bank"
}

func (plugin) InitContextData(context map[string]string) {
	Plugin.context = context
}

func (plugin) SupportsPlanPayment() bool {
	return false
}

func (plugin) ExecuteTransaction(data pspdto.TransactionDTO) (pspdto.TransactionCreatedDTO, error) {
	if data.PricingPlan {
		return pspdto.TransactionCreatedDTO{}, errors.New("bank does not support plan payment")
	}
	return transactions.PrepareTransaction(data, &Plugin.context)
}

func (plugin) CaptureTransaction(id string, plan bool) (bool, error) {
	if plan {
		return false, errors.New("bank does not allow for plan processing")
	}
	return transactions.CheckPaymentStatusSuccess(id)
}

var Plugin plugin

func main() {
	var ctx map[string]string = make(map[string]string, 0)
	ctx["000001"] = "https://remotehost.lol"
	Plugin.InitContextData(ctx)
	transactionCreatedDTO, err := Plugin.ExecuteTransaction(pspdto.TransactionDTO{
		PspTransactionId:            "1",
		OrderId:                     "o1",
		PayeeId:                     "0000011",
		PayeeSecret:                 "supersecret",
		Currency:                    "USD",
		Amount:                      "1",
		ClientBusinessName:          "ime",
		SuccessUrl:                  "success",
		FailUrl:                     "fail",
		ErrorUrl:                    "error",
		PricingPlan:                 false,
		PaymentInterval:             0,
		NumberOfInstallments:        0,
		InstallmentUnit:             "",
		InstallmentDelayedTimeUnits: 0,
	})
	if err != nil {
		return
	}
	fmt.Println(transactionCreatedDTO)
}
