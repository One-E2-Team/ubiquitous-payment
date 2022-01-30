package service

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"ubiquitous-payment/bank/dto"
	"ubiquitous-payment/bank/handler/mapper"
	"ubiquitous-payment/bank/model"
)

func (service *Service) Pay(issuerCard dto.IssuerCardDTO, paymentUrlId string) (*dto.PaymentResponseDTO, error) {
	if !strings.HasPrefix(issuerCard.Pan, os.Getenv("PAN_PREFIX")) {
		//TODO: call PCC
	}

	creditCard, err := service.BankRepository.GetCreditCard(issuerCard.Pan)
	if err != nil {
		return nil, err
	}

	if creditCard.Cvc != issuerCard.Cvc || creditCard.HolderName != issuerCard.HolderName ||
		creditCard.ValidUntil != issuerCard.ValidUntil {
		return nil, errors.New("bad credit card data")
	}

	transaction, err := service.BankRepository.GetTransactionByPaymentUrlId(paymentUrlId)
	if err != nil {
		return nil, err
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
