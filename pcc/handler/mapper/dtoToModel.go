package mapper

import (
	"ubiquitous-payment/pcc/dto"
	"ubiquitous-payment/pcc/model"
)

func PccOrderDtoToPccOrder(dto dto.PccOrderDto) *model.PccOrder{
	return &model.PccOrder{
		AcquirerTransactionId : dto.AcquirerTransactionId,
		AcquirerTimestamp: dto.AcquirerTimestamp,
		AcquirerPanPrefix: dto.AcquirerPanPrefix,
		Amount: dto.Amount,
		Currency: dto.Currency,
		IssuerPAN: dto.IssuerPAN,
		IssuerCVC: dto.IssuerCVC,
		IssuerValidUntil: dto.IssuerValidUntil,
		IssuerHolderName: dto.IssuerHolderName,
	}
}
