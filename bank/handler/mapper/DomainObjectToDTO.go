package mapper

import (
	"time"
	"ubiquitous-payment/bank/dto"
	"ubiquitous-payment/bank/model"
)

func TransactionToPaymentResponseDTO(transaction model.Transaction) *dto.PaymentResponseDTO {
	return &dto.PaymentResponseDTO{
		MerchantOrderId:   transaction.MerchantOrderID,
		AcquirerOrderId:   transaction.MerchantId,
		AcquirerTimestamp: time.Now(),
		PaymentId:         transaction.PaymentId,
		TransactionStatus: transaction.TransactionStatus,
	}
}

func PccOrderToPccOrderDTO(pccOrder model.PccOrder) *dto.PccOrderDTO {
	return &dto.PccOrderDTO{
		AcquirerTransactionId: pccOrder.AcquirerTransactionId,
		AcquirerTimestamp:     pccOrder.AcquirerTimestamp,
		AcquirerPanPrefix:     pccOrder.AcquirerPanPrefix,
		Amount:                pccOrder.Amount,
		Currency:              pccOrder.Currency,
		IssuerPAN:             pccOrder.IssuerPan,
		IssuerCVC:             pccOrder.IssuerCvc,
		IssuerValidUntil:      pccOrder.IssuerValidUntil,
		IssuerHolderName:      pccOrder.IssuerHolderName,
	}
}
