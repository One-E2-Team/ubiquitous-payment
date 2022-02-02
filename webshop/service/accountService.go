package service

import "ubiquitous-payment/webshop/model"

func (service *Service) GetAccountsByPaymentName(name string, loggedUserId uint) ([]model.Account, error) {
	paymentType, err := service.WSRepository.GetPaymentTypeByName(name)
	if err != nil{
		return nil, err
	}
	return service.WSRepository.GetAccountsByProfileIdAndPaymentType(loggedUserId, paymentType.ID)
}
