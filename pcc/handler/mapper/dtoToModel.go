package mapper

import (
	"ubiquitous-payment/pcc/dto"
	"ubiquitous-payment/pcc/model"
	"ubiquitous-payment/util"
)

func PccOrderDtoToPccOrder(dto dto.PccOrderDto) *model.PccOrder {
	return &model.PccOrder{
		AcquirerTransactionId: dto.AcquirerTransactionId,
		AcquirerTimestamp:     dto.AcquirerTimestamp,
		AcquirerPanPrefix:     dto.AcquirerPanPrefix,
		MerchantId:            dto.MerchantId,
		Amount:                dto.Amount,
		Currency:              dto.Currency,
		IssuerPAN:             util.GetEncryptedString(dto.IssuerPAN),
		IssuerCVC:             util.GetEncryptedString(dto.IssuerCVC),
		IssuerValidUntil:      dto.IssuerValidUntil,
		IssuerHolderName:      dto.IssuerHolderName,
	}
}
