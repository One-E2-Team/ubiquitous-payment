package service

import (
	"ubiquitous-payment/webshop/model"
)

func (service *Service) CreateProduct(product model.Product) error {
	// TODO: add picture in storage
	return service.WSRepository.CreateProduct(&product)
}

func (service *Service) UpdateProduct(productID uint, updatedProduct model.Product) error {
	product, err := service.WSRepository.GetProduct(productID)
	if err != nil {
		return err
	}
	if product.Price != updatedProduct.Price {
		product.Deactivate()
		err = service.WSRepository.UpdateProduct(product)
		if err != nil {
			return err
		}
		return service.WSRepository.CreateProduct(&updatedProduct)
	}
	product.Update(updatedProduct)
	return service.WSRepository.UpdateProduct(product)
}
