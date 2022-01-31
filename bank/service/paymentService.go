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

func (service *Service) Pay(issuerCard dto.IssuerCardDTO, paymentUrlId string) (*dto.PaymentResponseDTO, error) {
	transaction, err := service.BankRepository.GetTransactionByPaymentUrlId(paymentUrlId)
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(issuerCard.Pan, os.Getenv("PAN_PREFIX")) {
		err = service.proceedPaymentToPcc(issuerCard, transaction)
	}

	creditCard, err := service.BankRepository.GetCreditCard(issuerCard.Pan)
	if err != nil {
		return nil, err
	}

	if creditCard.Cvc != issuerCard.Cvc || creditCard.HolderName != issuerCard.HolderName ||
		creditCard.ValidUntil != issuerCard.ValidUntil {
		return nil, errors.New("bad credit card data")
	}

	err = service.payInSameBank(issuerCard.Pan, transaction)
	if err != nil {
		fmt.Println(err)
		transaction.TransactionStatus = model.FAILED
	} else {
		transaction.TransactionStatus = model.FULFILLED
	}
	err = service.BankRepository.Update(transaction)
	return mapper.TransactionToPaymentResponseDTO(*transaction), err
}

func (service *Service) payInSameBank(issuerPan string, transaction *model.Transaction) error {
	issuerAccount, err := service.BankRepository.GetClientAccountByPan(issuerPan)
	if err != nil {
		return err
	}

	acquirerAccount, err := service.BankRepository.GetClientAccount(transaction.MerchantId)
	if err != nil {
		return err
	}

	//TODO: check currency
	if issuerAccount.Amount < transaction.AmountRsd {
		return errors.New("not enough money")
	}

	issuerAccount.Amount -= transaction.AmountRsd
	acquirerAccount.Amount += transaction.AmountRsd
	err = service.BankRepository.Update(issuerAccount)
	if err != nil {
		return err
	}
	return service.BankRepository.Update(acquirerAccount)
}

func (service *Service) proceedPaymentToPcc(issuerCard dto.IssuerCardDTO, transaction *model.Transaction) error {
	pccOrder := model.PccOrder{
		AcquirerTransactionId: transaction.ID,
		AcquirerTimestamp:     time.Now(),
		AcquirerPanPrefix:     os.Getenv("PAN_PREFIX"),
		Amount:                transaction.AmountRsd, //TODO: currency
		Currency:              transaction.Currency,
		IssuerPan:             issuerCard.Pan,
		IssuerCvc:             issuerCard.Cvc,
		IssuerValidUntil:      issuerCard.ValidUntil,
		IssuerHolderName:      issuerCard.HolderName,
	}

	err := service.BankRepository.Update(&pccOrder)
	if err != nil {
		return err
	}

	client := &http.Client{}
	pccHost, pccPort := util.GetPccHostAndPort()
	jsonReq, _ := json.Marshal(pccOrder)
	req, err := http.NewRequest(http.MethodPost, util.GetPccProtocol()+"://"+pccHost+":"+pccPort+"/pcc-order", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println(err)
		return err
	}
	resp, err := client.Do(req)
	// TODO: implement pcc response
	fmt.Println(resp)
	return nil
}
