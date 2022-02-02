package repository

import "ubiquitous-payment/webshop/model"

func (repo *Repository) GetValidPaymentTypes() ([]model.PaymentType, error) {
	var paymentTypes []model.PaymentType
	result := repo.RelationalDatabase.Table("payment_types").Where("is_valid=1").Find(&paymentTypes)
	return paymentTypes, result.Error
}

func (repo *Repository) GetPaymentTypeByName(name string) (model.PaymentType, error) {
	var paymentType model.PaymentType
	result := repo.RelationalDatabase.Table("payment_types").Where("name=?", name).First(&paymentType)
	return paymentType, result.Error
}
