package service

import (
	"errors"
	"github.com/google/uuid"
	"ubiquitous-payment/bank/dto"
	"ubiquitous-payment/bank/model"
	"ubiquitous-payment/util"
)

func (service *Service) PspRequest(transaction model.Transaction, paymentMethod string) (*dto.PspResponseDTO, error) {
	clientAccount, err := service.Repository.GetClientAccount(transaction.MerchantId)
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
	err = service.Repository.CreateTransaction(&transaction)
	if err != nil {
		return nil, err
	}

	bankFrontHost, bankFrontPort := util.GetBankFrontHostAndPort()
	bankHost, bankPort := util.GetExternalBankHostAndPort()
	bankProtocol := util.GetBankProtocol()
	payTransactionUrl := ""
	if paymentMethod == "bank" {
		payTransactionUrl = bankProtocol + "://" + bankFrontHost + ":" + bankFrontPort + "/#/payment?id=" + transaction.PaymentUrlId
	} else if paymentMethod == "qrcode" {
		payTransactionUrl = bankProtocol + "://" + bankHost + ":" + bankPort + "/api/pay/" + transaction.PaymentUrlId
	} else {
		return nil, errors.New("request for unknown payment method in bank: " + paymentMethod)
	}
	paymentCheckUrl := bankProtocol + "://" + bankHost + ":" + bankPort + "/api/payment-check/{id}"
	return &dto.PspResponseDTO{
		PaymentId: transaction.PaymentId, PaymentUrl: payTransactionUrl, PaymentCheckUrl: paymentCheckUrl}, nil
}

func (service *Service) CheckPaymentStatus(id string) (*dto.PaymentResponseDTO, error) {
	transaction, err := service.Repository.GetTransactionByPaymentId(id)
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
