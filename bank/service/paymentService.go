package service

import (
	"bytes"
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

func (service *Service) Pay(issuerCard dto.IssuerCardDTO, paymentUrlId string) (string, error) {
	transaction, err := service.Repository.GetTransactionByPaymentUrlId(paymentUrlId)
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(issuerCard.Pan, os.Getenv("PAN_PREFIX")) {
		return service.proceedPaymentToPcc(issuerCard, transaction)
	}

	if !service.IsCreditCardDataValid(issuerCard) {
		return "", errors.New("bad credit card data")
	}

	err = service.payInSameBank(issuerCard.Pan, transaction)
	if err != nil {
		fmt.Println(err)
		transaction.TransactionStatus = model.FAILED
	} else {
		transaction.TransactionStatus = model.FULFILLED
	}
	err = service.Repository.Update(transaction)
	return transaction.GetURLByStatus(), err
}

func (service *Service) IssuerPay(pccOrderDto dto.PccOrderDTO) (*dto.PccResponseDTO, error) {
	transaction, err := mapper.PccOrderDTOToTransaction(pccOrderDto)
	if err != nil {
		return nil, err
	}
	if !service.IsCreditCardDataValid(mapper.PccOrderDtoToIssuerCardDto(pccOrderDto)) {
		return nil, errors.New("bad credit card data")
	}

	issuerAccount, err := service.Repository.GetClientAccountByPan(pccOrderDto.IssuerPAN)
	if err != nil {
		return nil, err
	}

	if issuerAccount.Amount < transaction.AmountRsd {
		return nil, errors.New("not enough money")
	}

	issuerAccount.Amount -= transaction.AmountRsd
	err = service.Repository.Update(issuerAccount)
	if err != nil {
		return nil, err
	}

	transaction.TransactionStatus = model.FULFILLED
	err = service.Repository.CreateTransaction(&transaction)
	if err != nil {
		return nil, err
	}
	return mapper.TransactionToPccResponseDTO(transaction), nil
}

func (service *Service) payInSameBank(issuerPan string, transaction *model.Transaction) error {
	issuerAccount, err := service.Repository.GetClientAccountByPan(issuerPan)
	if err != nil {
		return err
	}

	acquirerAccount, err := service.Repository.GetClientAccount(transaction.MerchantId)
	if err != nil {
		return err
	}

	if issuerAccount.Amount < transaction.AmountRsd {
		return errors.New("not enough money")
	}

	issuerAccount.Amount -= transaction.AmountRsd
	acquirerAccount.Amount += transaction.AmountRsd
	err = service.Repository.Update(issuerAccount)
	if err != nil {
		return err
	}
	return service.Repository.Update(acquirerAccount)
}

func (service *Service) proceedPaymentToPcc(issuerCard dto.IssuerCardDTO, transaction *model.Transaction) (string, error) {
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

	client := &http.Client{}
	pccHost, pccPort := util.GetPccHostAndPort()
	jsonReq, _ := json.Marshal(pccOrder)
	req, err := http.NewRequest(http.MethodPost, util.GetPccProtocol()+"://"+pccHost+":"+pccPort+"/pcc-order", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	resp, err := client.Do(req)

	var respDto dto.PccResponseDTO
	err = util.UnmarshalResponse(resp, &respDto)
	if err != nil {
		return "", err
	}

	transaction.TransactionStatus = respDto.OrderStatus
	err = service.Repository.Update(transaction)
	if err != nil {
		return "", err
	}

	acquirerAccount, err := service.Repository.GetClientAccount(transaction.MerchantId)
	if err != nil {
		return "", err
	}

	acquirerAccount.Amount += transaction.AmountRsd
	return transaction.GetURLByStatus(), service.Repository.Update(acquirerAccount)
}

func (service *Service) IsCreditCardDataValid(issuerCard dto.IssuerCardDTO) bool {
	creditCard, err := service.Repository.GetCreditCard(issuerCard.Pan)
	if err != nil {
		return false
	}

	return creditCard.Cvc == issuerCard.Cvc && creditCard.HolderName == issuerCard.HolderName &&
		creditCard.ValidUntil == issuerCard.ValidUntil //TODO: check valid until
}
