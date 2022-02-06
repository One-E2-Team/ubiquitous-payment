package mapper

import (
	"ubiquitous-payment/util"
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

func UpdateAccountDTOToAccount(accountDto dto.UpdateAccountDTO, loggedMerchantId uint) *model.Account {
	return &model.Account{
		AccountID:     util.GetEncryptedString(accountDto.AccountID),
		Secret:        util.GetEncryptedString(accountDto.Secret),
		PaymentTypeId: accountDto.PaymentTypeId,
		ProfileId:     loggedMerchantId,
	}
}
