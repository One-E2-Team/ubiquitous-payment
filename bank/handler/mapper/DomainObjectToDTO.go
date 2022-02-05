package mapper

import (
	"time"
	"ubiquitous-payment/bank/bankutil"
	"ubiquitous-payment/bank/dto"
	"ubiquitous-payment/bank/model"
	"ubiquitous-payment/util"
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

func AccountToAccountResponseDTO(account *model.ClientAccount) *dto.AccountResponseDTO {
	if account == nil {
		return nil
	}
	accountCards := account.CreditCards
	creditCards := make([]dto.CreditCardResponseDTO, len(accountCards))
	for i := 0; i < len(accountCards); i++ {
		creditCards[i] = CreditCardToCreditCardResponseDTO(accountCards[i])
	}

	return &dto.AccountResponseDTO{
		AccountNumber: account.AccountNumber.Data,
		Amount:        account.Amount,
		Secret:        account.Secret.Data,
		CreditCards:   creditCards,
	}
}

func TransactionToTransactionResponseDTO(transaction model.Transaction, amountPrefix string) dto.TransactionResponseDTO {
	return dto.TransactionResponseDTO{
		Amount:                amountPrefix + util.Float32ToString(transaction.Amount),
		Currency:              transaction.Currency,
		AcquirerAccountNumber: bankutil.CensorPaymentString(transaction.MerchantId),
		IssuerPan:             bankutil.CensorPaymentString(transaction.IssuerPan),
		Timestamp:             transaction.MerchantTimestamp,
		TransactionStatus:     string(transaction.TransactionStatus),
	}
}

func CreditCardToCreditCardResponseDTO(creditCard model.CreditCard) dto.CreditCardResponseDTO {
	return dto.CreditCardResponseDTO{
		Pan:        creditCard.Pan.Data,
		Cvc:        creditCard.Cvc.Data,
		HolderName: creditCard.HolderName,
		ValidUntil: creditCard.ValidUntil,
	}
}
