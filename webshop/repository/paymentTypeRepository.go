package repository

import "ubiquitous-payment/webshop/model"

func (repo *Repository) GetValidPaymentTypes() ([]model.PaymentType, error) {
	var paymentTypes []model.PaymentType
	result := repo.RelationalDatabase.Table("payment_types").Where("is_valid=1").Find(&paymentTypes)
	return paymentTypes, result.Error
}
