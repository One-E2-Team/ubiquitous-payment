package service

import "ubiquitous-payment/pcc/model"

func (service *Service) GetBankByPanPrefix(panPrefix string) (*model.Bank, error){
	return service.Repository.GetBankByPanPrefix(panPrefix)
}
