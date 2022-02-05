package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"ubiquitous-payment/bank/dto"
	"ubiquitous-payment/bank/handler/mapper"
	"ubiquitous-payment/bank/model"
	"ubiquitous-payment/util"
)

var loggingService = "bank_" + os.Getenv("PAN_PREFIX")

func (service *Service) Pay(issuerCard dto.IssuerCardDTO, paymentUrlId string) string {
	transaction, err := service.Repository.GetTransactionByPaymentUrlId(paymentUrlId)
	if err != nil {
		util.Logging(util.ERROR, "Service.Pay", err.Error(), loggingService)
		return ""
	}

	if !strings.HasPrefix(issuerCard.Pan, os.Getenv("PAN_PREFIX")) {
		return service.proceedPaymentToPcc(issuerCard, transaction)
	}

	if !service.isCreditCardDataValid(issuerCard) {
		util.Logging(util.ERROR, "Service.Pay", "bad issuer credit card data", loggingService)
		transaction = service.saveTransactionStatus(transaction, model.FAILED)
		return transaction.GetURLByStatus()
	}

	return service.payInSameBank(issuerCard.Pan, transaction)
}

func (service *Service) IssuerPay(pccOrderDto dto.PccOrderDTO) (*dto.PccResponseDTO, error) {
	transaction, err := mapper.PccOrderDTOToTransaction(pccOrderDto)
	if err != nil {
		util.Logging(util.ERROR, "Service.IssuerPay", err.Error(), loggingService)
		transaction = service.saveTransactionStatus(transaction, model.ERROR)
		return nil, err
	}
	if !service.isCreditCardDataValid(mapper.PccOrderDtoToIssuerCardDto(pccOrderDto)) {
		util.Logging(util.ERROR, "Service.IssuerPay", "bad issuer credit card data", loggingService)
		transaction = service.saveTransactionStatus(transaction, model.FAILED)
		return nil, errors.New("bad issuer credit card data")
	}

	issuerAccount, err := service.Repository.GetClientAccountByPan(pccOrderDto.IssuerPAN)
	if err != nil {
		util.Logging(util.ERROR, "Service.IssuerPay", err.Error(), loggingService)
		transaction = service.saveTransactionStatus(transaction, model.ERROR)
		return nil, err
	}

	if issuerAccount.Amount < transaction.AmountRsd {
		util.Logging(util.ERROR, "Service.IssuerPay", "not enough money on issuer's account", loggingService)
		transaction = service.saveTransactionStatus(transaction, model.FAILED)
		return nil, errors.New("not enough money on issuer's account")
	}

	issuerAccount.Amount -= transaction.AmountRsd
	err = service.Repository.Update(issuerAccount)
	if err != nil {
		util.Logging(util.ERROR, "Service.IssuerPay", err.Error(), loggingService)
		transaction = service.saveTransactionStatus(transaction, model.ERROR)
		return nil, err
	}

	transaction = service.saveTransactionStatus(transaction, model.FULFILLED)
	return mapper.TransactionToPccResponseDTO(*transaction), nil
}

func (service *Service) payInSameBank(issuerPan string, transaction *model.Transaction) string {
	issuerAccount, err := service.Repository.GetClientAccountByPan(issuerPan)
	if err != nil {
		util.Logging(util.ERROR, "Service.payInSameBank", err.Error(), loggingService)
		transaction = service.saveTransactionStatus(transaction, model.ERROR)
		return transaction.GetURLByStatus()
	}

	acquirerAccount, err := service.Repository.GetClientAccount(transaction.MerchantId)
	if err != nil {
		util.Logging(util.ERROR, "Service.payInSameBank", err.Error(), loggingService)
		transaction = service.saveTransactionStatus(transaction, model.ERROR)
		return transaction.GetURLByStatus()
	}

	if issuerAccount.Amount < transaction.AmountRsd {
		util.Logging(util.ERROR, "Service.payInSameBank", "not enough money on issuer's account", loggingService)
		transaction = service.saveTransactionStatus(transaction, model.FAILED)
		return transaction.GetURLByStatus()
	}

	issuerAccount.Amount -= transaction.AmountRsd
	acquirerAccount.Amount += transaction.AmountRsd
	err = service.Repository.Update(issuerAccount)
	if err != nil {
		util.Logging(util.ERROR, "Service.payInSameBank", err.Error(), loggingService)
		transaction = service.saveTransactionStatus(transaction, model.ERROR)
		return transaction.GetURLByStatus()
	}
	err = service.Repository.Update(acquirerAccount)
	if err != nil {
		util.Logging(util.ERROR, "Service.payInSameBank", err.Error(), loggingService)
		transaction = service.saveTransactionStatus(transaction, model.ERROR)
	} else {
		transaction = service.saveTransactionStatus(transaction, model.FULFILLED)
	}
	return transaction.GetURLByStatus()
}

func (service *Service) proceedPaymentToPcc(issuerCard dto.IssuerCardDTO, transaction *model.Transaction) string {
	pccOrder := dto.PccOrderDTO{
		AcquirerTransactionId: transaction.ID,
		AcquirerTimestamp:     time.Now(),
		AcquirerPanPrefix:     os.Getenv("PAN_PREFIX"),
		MerchantId:            transaction.MerchantId,
		Amount:                transaction.AmountRsd,
		Currency:              "RSD",
		IssuerPAN:             issuerCard.Pan,
		IssuerCVC:             issuerCard.Cvc,
		IssuerValidUntil:      issuerCard.ValidUntil,
		IssuerHolderName:      issuerCard.HolderName,
	}

	jsonReq, _ := json.Marshal(pccOrder)
	pccHost, pccPort := util.GetExternalPccHostAndPort()
	resp, err := util.CrossServiceRequest(http.MethodPost, util.GetPccProtocol()+"://"+pccHost+":"+pccPort+"/pcc-order", jsonReq, nil)
	if err != nil {
		util.Logging(util.ERROR, "Service.proceedPaymentToPcc", err.Error(), loggingService)
		transaction = service.saveTransactionStatus(transaction, model.ERROR)
		return transaction.GetURLByStatus()
	}

	var respDto dto.PccResponseDTO
	err = util.UnmarshalResponse(resp, &respDto)
	if err != nil {
		util.Logging(util.ERROR, "Service.proceedPaymentToPcc", err.Error(), loggingService)
		transaction = service.saveTransactionStatus(transaction, model.ERROR)
		return transaction.GetURLByStatus()
	}

	transaction.TransactionStatus = respDto.OrderStatus
	err = service.Repository.Update(transaction)
	if err != nil {
		util.Logging(util.ERROR, "Service.proceedPaymentToPcc", err.Error(), loggingService)
		transaction = service.saveTransactionStatus(transaction, model.ERROR)
		return transaction.GetURLByStatus()
	}

	acquirerAccount, err := service.Repository.GetClientAccount(transaction.MerchantId)
	if err != nil {
		util.Logging(util.ERROR, "Service.proceedPaymentToPcc", err.Error(), loggingService)
		transaction = service.saveTransactionStatus(transaction, model.ERROR)
		return transaction.GetURLByStatus()
	}

	acquirerAccount.Amount += transaction.AmountRsd
	err = service.Repository.Update(acquirerAccount)
	if err != nil {
		util.Logging(util.ERROR, "Service.proceedPaymentToPcc", err.Error(), loggingService)
		transaction = service.saveTransactionStatus(transaction, model.ERROR)
	} else {
		transaction = service.saveTransactionStatus(transaction, model.FULFILLED)
	}
	return transaction.GetURLByStatus()
}

func (service *Service) isCreditCardDataValid(issuerCard dto.IssuerCardDTO) bool {
	creditCard, err := service.Repository.GetCreditCard(issuerCard.Pan)
	if err != nil {
		return false
	}

	return creditCard.Cvc == issuerCard.Cvc && creditCard.HolderName == issuerCard.HolderName &&
		creditCard.ValidUntil == issuerCard.ValidUntil //TODO: check valid until
}

func (service *Service) saveTransactionStatus(transaction *model.Transaction, status model.TransactionStatus) *model.Transaction {
	transaction.TransactionStatus = status

	if transaction.ID == 0 {
		_ = service.Repository.CreateTransaction(transaction)
		return transaction
	}

	err := service.Repository.Update(transaction)
	if err != nil {
		fmt.Println(err)
	}
	return transaction
}
