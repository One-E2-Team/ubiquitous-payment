package mapper

import (
	"ubiquitous-payment/webshop/dto"
	"ubiquitous-payment/webshop/model"
)

func ProductDTOToProduct(productDto dto.ProductDTO, loggedMerchantId uint) model.Product {
	return model.Product{
		Name:                productDto.Name,
		Price:               productDto.Price,
		Currency:            productDto.Currency,
		Description:         productDto.Description,
		IsActive:            true,
		NumOfInstallments:   productDto.NumOfInstallments,
		MerchantProfileId:   loggedMerchantId,
		RecurringType:       productDto.RecurringType,
		DelayedInstallments: productDto.DelayedInstallments,
	}
}
