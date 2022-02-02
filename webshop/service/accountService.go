package service

import (
	"errors"
	"strings"
	"ubiquitous-payment/webshop/model"
)

func (service *Service) GetAccountsByPaymentName(name string, loggedUserId uint) ([]model.Account, error) {
	paymentType, err := service.WSRepository.GetPaymentTypeByName(name)
	if err != nil{
		return nil, err
	}
	return service.WSRepository.GetAccountsByProfileIdAndPaymentType(loggedUserId, paymentType.ID)
}

func (service *Service) UpdateAccount(newAcc *model.Account, id uint) error {
	 acc := service.WSRepository.GetAccountById(id)
	 if acc == nil{
	 	return errors.New("account does not exist")
	 }
	 acc.AccountID = newAcc.AccountID
	 acc.Secret = newAcc.Secret
	 return service.WSRepository.Update(acc)
}

func (service *Service) CreateAccount(account *model.Account) error {
	account.AccountID = strings.Trim(account.AccountID, "\t")
	account.Secret = strings.Trim(account.Secret, "\t")
	return service.WSRepository.CreateAccount(account)
}

func (service *Service) DeleteAccount(id uint) error {
	return service.WSRepository.DeleteAccount(id)
}
