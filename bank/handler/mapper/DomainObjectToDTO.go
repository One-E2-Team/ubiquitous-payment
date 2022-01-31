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

func TransactionToPccResponseDTO(transaction model.Transaction) *dto.PccResponseDTO {
	return &dto.PccResponseDTO{
		IssuerOrderId:   transaction.ID,
		IssuerTimestamp: time.Now(),
		OrderStatus:     transaction.TransactionStatus,
	}
}
