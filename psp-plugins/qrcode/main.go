package main

import (
	"errors"
	"ubiquitous-payment/psp-plugins/bank/transactions"
	"ubiquitous-payment/psp-plugins/pspdto"
)

type plugin struct {
	context map[string]string
}

func (plugin) Name() string {
	return "qrcode"
}

func (plugin) InitContextData(context map[string]string) {
	Plugin.context = context
}

func (plugin) SupportsPlanPayment() bool {
	return false
}

func (plugin) ExecuteTransaction(data pspdto.TransactionDTO) (pspdto.TransactionCreatedDTO, error) {
	if data.PricingPlan {
		return pspdto.TransactionCreatedDTO{}, errors.New("qrcode does not support plan payment")
	}

	return transactions.PrepareTransaction(data, &Plugin.context, Plugin.Name())
}

func (plugin) CaptureTransaction(id string, plan bool) (bool, error) {
	if plan {
		return false, errors.New("qrcode does not allow for plan processing")
	}
	return transactions.CheckPaymentStatusSuccess(id)
}

var Plugin plugin
