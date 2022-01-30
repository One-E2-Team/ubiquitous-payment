package service

import (
	"fmt"
	"github.com/google/uuid"
	"ubiquitous-payment/bank/dto"
	"ubiquitous-payment/bank/model"
	"ubiquitous-payment/util"
)

func (service *Service) PspRequest(transaction model.Transaction) (*dto.PspResponseDTO, error) {
	clientAccount, err := service.BankRepository.GetClientAccount(transaction.MerchantId)
	if err != nil {
		return nil, err
	}

	if clientAccount.Secret != transaction.MerchantPassword {
		return nil, fmt.Errorf("bad merchant credentials")
	}

	transaction.PaymentId = uuid.NewString()
	transaction.PaymentUrlId = uuid.NewString()
	err = service.BankRepository.CreateTransaction(&transaction)
	if err != nil {
		return nil, err
	}

	bankHost, bankPort := util.GetBankHostAndPort()
	payTransactionUrl := "https://" + bankHost + ":" + bankPort + "/api/pay/" + transaction.PaymentUrlId
	paymentCheckUrl := "https://" + bankHost + ":" + bankPort + "/api/payment-check/{id}"
	return &dto.PspResponseDTO{
		PaymentId: transaction.PaymentId, PaymentUrl: payTransactionUrl, PaymentCheckUrl: paymentCheckUrl}, nil
}
