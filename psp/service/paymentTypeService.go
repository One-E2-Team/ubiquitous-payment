package service

import "ubiquitous-payment/psp/model"

func (service *Service) paymentTypeListContains(list []model.PaymentType, paymentTypeName string) bool {
	for _, pt := range list {
		if pt.Name == paymentTypeName {
			return true
		}
	}
	return false
}
