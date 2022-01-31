package mapper

import (
	"time"
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
		MerchantTimestamp: time.Now(), //todo: change
	}
}

func PccOrderDTOToTransaction(pccOrderDto dto.PccOrderDTO) model.Transaction {
	return model.Transaction{
		Amount:            pccOrderDto.Amount,
		Currency:          pccOrderDto.Currency,
		MerchantId:        pccOrderDto.MerchantId,
		IssuerPan:         pccOrderDto.IssuerPAN,
		MerchantTimestamp: time.Now(),
	}
}
