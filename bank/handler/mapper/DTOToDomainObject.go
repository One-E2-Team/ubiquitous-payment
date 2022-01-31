package mapper

import (
	"time"
	"ubiquitous-payment/bank/dto"
	"ubiquitous-payment/bank/model"
	"ubiquitous-payment/bank/util"
)

func PspRequestDTOToTransaction(pspRequest dto.PspRequestDTO) (model.Transaction, error) {
	amountRsd := pspRequest.Amount
	var err error
	if pspRequest.Currency != "RSD" {
		amountRsd, err = util.CurrencyConversion(pspRequest.Currency, "RSD", pspRequest.Amount)
		if err != nil {
			return model.Transaction{}, err
		}
	}
	return model.Transaction{
		Amount:            pspRequest.Amount,
		AmountRsd:         amountRsd,
		Currency:          pspRequest.Currency,
		MerchantId:        pspRequest.MerchantId,
		MerchantPassword:  pspRequest.MerchantPassword,
		SuccessURL:        pspRequest.SuccessURL,
		FailURL:           pspRequest.FailURL,
		ErrorURL:          pspRequest.ErrorURL,
		MerchantOrderID:   pspRequest.MerchantOrderID,
		MerchantTimestamp: time.Now(), //todo: change
	}, nil
}

func PccOrderDTOToTransaction(pccOrderDto dto.PccOrderDTO) (model.Transaction, error) {
	amountRsd := pccOrderDto.Amount
	var err error
	if pccOrderDto.Currency != "RSD" {
		amountRsd, err = util.CurrencyConversion(pccOrderDto.Currency, "RSD", pccOrderDto.Amount)
		if err != nil {
			return model.Transaction{}, err
		}
	}
	return model.Transaction{
		Amount:            pccOrderDto.Amount,
		AmountRsd:         amountRsd,
		Currency:          pccOrderDto.Currency,
		MerchantId:        pccOrderDto.MerchantId,
		IssuerPan:         pccOrderDto.IssuerPAN,
		MerchantTimestamp: time.Now(),
	}, nil
}
