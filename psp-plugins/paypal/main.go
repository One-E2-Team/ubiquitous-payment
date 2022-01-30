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

func (p plugin) InitContextData(_ map[string]string) {
	return
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

func (p plugin) CaptureTransaction(id string, plan bool) (bool, error) {
	if plan {
		return transactions.CaptureSubscriptionApproval(id)
	}
	return transactions.CaptureOrderPayment(id)
}

var Plugin plugin

func main() {
	ret, err := Plugin.ExecuteTransaction(pspdto.TransactionDTO{
		PspTransactionId:            "T-0005",
		OrderId:                     "O-0005",
		PayeeId:                     "sb-064747x8893734@business.example.com",
		PayeeSecret:                 "35AYF8PFJWGPS",
		Currency:                    "USD",
		Amount:                      "99.90",
		ClientBusinessName:          "PORNJAVA.COM",
		SuccessUrl:                  "https://www.igorsikuljak.rs/success",
		FailUrl:                     "https://www.igorsikuljak.rs/fail",
		ErrorUrl:                    "https://www.igorsikuljak.rs/error",
		PricingPlan:                 true,
		PaymentInterval:             1,
		NumberOfInstallments:        0,
		InstallmentUnit:             pspdto.Month,
		InstallmentDelayedTimeUnits: 0, //3,
	})
	fmt.Println(ret, err)
}
