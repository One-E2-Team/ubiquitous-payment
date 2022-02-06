package mapper

import (
	"ubiquitous-payment/pcc/dto"
	"ubiquitous-payment/pcc/model"
)

func PccOrderToIssuerBankRequestDTO(pccOrder model.PccOrder) *dto.IssuerBankRequestDTO {
	return &dto.IssuerBankRequestDTO{
		AcquirerTransactionId: pccOrder.AcquirerTransactionId,
		AcquirerTimestamp:     pccOrder.AcquirerTimestamp,
		AcquirerPanPrefix:     pccOrder.AcquirerPanPrefix,
		MerchantId:            pccOrder.MerchantId,
		Amount:                pccOrder.Amount,
		Currency:              pccOrder.Currency,
		IssuerPAN:             pccOrder.IssuerPAN.Data,
		IssuerCVC:             pccOrder.IssuerCVC.Data,
		IssuerValidUntil:      pccOrder.IssuerValidUntil,
		IssuerHolderName:      pccOrder.IssuerHolderName,
		IssuerOrderId:         pccOrder.IssuerOrderId,
		IssuerTimestamp:       pccOrder.IssuerTimestamp,
		OrderStatus:           pccOrder.OrderStatus,
	}
}
