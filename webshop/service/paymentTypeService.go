package service

import "ubiquitous-payment/webshop/model"

func (service *Service) GetValidPaymentTypes() ([]model.PaymentType, error) {
	return service.WSRepository.GetValidPaymentTypes()
}
