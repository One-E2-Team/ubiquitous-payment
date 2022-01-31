package mapper

import (
	"ubiquitous-payment/bank/dto"
	"ubiquitous-payment/bank/model"
)

func PspRequestDTOToTransaction(pspRequest dto.PspRequestDTO) model.Transaction {
	return model.Transaction{
		AmountRsd:         pspRequest.Amount, //TODO: check currency
		Currency:          pspRequest.Currency,
		MerchantId:        pspRequest.MerchantId,
		MerchantPassword:  pspRequest.MerchantPassword,
		SuccessURL:        pspRequest.SuccessURL,
		FailURL:           pspRequest.FailURL,
		ErrorURL:          pspRequest.ErrorURL,
		MerchantOrderID:   pspRequest.MerchantOrderID,
		MerchantTimestamp: pspRequest.MerchantTimestamp,
	}
}
