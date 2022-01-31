package service

import (
	"errors"
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
		return nil, errors.New("bad merchant credentials")
	}

	transaction.PaymentId = uuid.NewString()
	transaction.SuccessURL += "?token=" + transaction.PaymentId
	transaction.FailURL += "?token=" + transaction.PaymentId
	transaction.ErrorURL += "?token=" + transaction.PaymentId
	transaction.PaymentUrlId = uuid.NewString()
	err = service.BankRepository.CreateTransaction(&transaction)
	if err != nil {
		return nil, err
	}

	bankHost, bankPort := util.GetBankHostAndPort()
	bankProtocol := util.GetBankProtocol()
	payTransactionUrl := bankProtocol + "://" + bankHost + ":" + bankPort + "/api/pay/" + transaction.PaymentUrlId
	paymentCheckUrl := bankProtocol + "://" + bankHost + ":" + bankPort + "/api/payment-check/{id}"
	return &dto.PspResponseDTO{
		PaymentId: transaction.PaymentId, PaymentUrl: payTransactionUrl, PaymentCheckUrl: paymentCheckUrl}, nil
}

func (service *Service) CheckPaymentStatus(id string) (*dto.PaymentResponseDTO, error) {
	transaction, err := service.BankRepository.GetTransactionByPaymentId(id)
	if err != nil {
		return nil, err
	}
	return &dto.PaymentResponseDTO{
		MerchantOrderId:   transaction.MerchantOrderID,
		AcquirerOrderId:   transaction.MerchantId,
		AcquirerTimestamp: transaction.MerchantTimestamp,
		PaymentId:         transaction.PaymentId,
		TransactionStatus: transaction.TransactionStatus,
	}, nil
}
