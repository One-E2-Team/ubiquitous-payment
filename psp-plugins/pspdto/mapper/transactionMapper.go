package mapper

import (
	"strconv"
	"ubiquitous-payment/psp-plugins/pspdto"
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/psputil"
	"ubiquitous-payment/util"
)

func TransactionToTransactionDTO(transaction model.Transaction, plugin psputil.Plugin) (pspdto.TransactionDTO, error) {
	account, err := transaction.GetSelectedAccount()
	if err != nil {
		return pspdto.TransactionDTO{}, err
	}
	pspHost, pspPort := util.GetPSPHostAndPort()
	initialUrl := util.GetPSPProtocol() + "://" + pspHost + ":" + pspPort + "/api/psp"
	pricingPlan, err := transaction.IsPricingPlan(plugin)
	if err != nil {
		return pspdto.TransactionDTO{}, err
	}

	return pspdto.TransactionDTO{
		PspTransactionId:            transaction.PSPId,
		OrderId:                     transaction.MerchantOrderID,
		PayeeId:                     account.AccountID,
		PayeeSecret:                 account.Secret,
		Currency:                    transaction.Currency,
		Amount:                      strconv.FormatFloat(float64(transaction.Amount), 'f', 2, 64),
		ClientBusinessName:          transaction.WebShopID,
		SuccessUrl:                  initialUrl + "/payment-success",
		FailUrl:                     initialUrl + "/payment-fail",
		ErrorUrl:                    transaction.ErrorURL,
		PricingPlan:                 pricingPlan,
		PaymentInterval:             1,
		NumberOfInstallments:        int(transaction.Recurring.InstallmentCount),
		InstallmentUnit:             model.GetInstallmentUnitByRecurringType(transaction.Recurring.Type),
		InstallmentDelayedTimeUnits: int(transaction.Recurring.DelayedInstallmentCount),
	}, nil

}
