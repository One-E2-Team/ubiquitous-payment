package mapper

import (
	"strconv"
	"ubiquitous-payment/psp-plugins/pspdto"
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/psputil"
)

func TransactionToTransactionDTO(transaction model.Transaction, plugin psputil.Plugin) (pspdto.TransactionDTO, error) {
	account, err := transaction.GetSelectedAccount()
	if err != nil {
		return pspdto.TransactionDTO{}, err
	}
	//pspHost, pspPort := util.GetPSPHostAndPort()
	initialUrl := "http" + "://" + "igorsikuljak.rs" + "/api/psp"
	pricingPlan, err := transaction.IsPricingPlan(plugin)
	if err != nil {
		return pspdto.TransactionDTO{}, err
	}

	numberOfInstallments := 1
	installmentUnit := pspdto.Month
	installmentDelayedTimeUnits := 0
	if transaction.Recurring != nil {
		numberOfInstallments = int(transaction.Recurring.InstallmentCount)
		installmentUnit = model.GetInstallmentUnitByRecurringType(transaction.Recurring.Type)
		installmentDelayedTimeUnits = int(transaction.Recurring.DelayedInstallmentCount)
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
		NumberOfInstallments:        numberOfInstallments,
		InstallmentUnit:             installmentUnit,
		InstallmentDelayedTimeUnits: installmentDelayedTimeUnits,
	}, nil

}
