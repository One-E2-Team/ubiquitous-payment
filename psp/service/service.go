package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ubiquitous-payment/psp/dto"
	"ubiquitous-payment/psp/repository"
)

type Service struct {
	PSPRepository *repository.Repository
}

func (service *Service) GetDataForQrCode(id string) (dto.QrCodeDTO, error) {
	ret := dto.QrCodeDTO{}
	primitiveId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return ret, err
	}
	t, err := service.PSPRepository.GetTransactionById(primitiveId)
	if err != nil{
		return ret, err
	}
	for _, acc := range t.MerchantAccounts{
		if acc.PaymentType.Name == "qrcode" {
			ret.AccountID = acc.AccountID
			break
		}
	}
	ret.WebShopName = t.WebShopID
	ret.Amount = fmt.Sprintf("%f", t.Amount)
	ret.Currency = t.Currency

	return ret, nil
}
