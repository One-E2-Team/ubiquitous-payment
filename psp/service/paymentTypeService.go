package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ubiquitous-payment/psp/dto"
	"ubiquitous-payment/psp/model"
)

func (service *Service) paymentTypeListContains(list []model.PaymentType, paymentTypeName string) bool {
	for _, pt := range list {
		if pt.Name == paymentTypeName {
			return true
		}
	}
	return false
}

func (service *Service) GetMyPaymentTypes(profileId string) (dto.MyPaymentTypesDTO, error) {
	var ret dto.MyPaymentTypesDTO
	retPaymentTypes := make([]string, 0)
	retMyPaymentTypes := make([]string, 0)
	allPaymentTypes, err := service.PSPRepository.GetAllPaymentTypes()
	if err != nil{
		return ret, err
	}
	for _, pt := range allPaymentTypes{
		retPaymentTypes = append(retPaymentTypes, pt.Name)
	}

	profileIdPrimitive, err := primitive.ObjectIDFromHex(profileId)
	if err != nil{
		return ret, err
	}
	user, err := service.PSPRepository.GetUserByID(profileIdPrimitive)
	if err != nil{
		return ret, err
	}
	webShopIdPrimitive, err := primitive.ObjectIDFromHex(user.WebShopId)
	if err != nil{
		return ret, err
	}
	webShop, err := service.PSPRepository.GetWebShopByID(webShopIdPrimitive)
	if err != nil{
		return ret, err
	}
	if len(webShop.PaymentTypes) > 0{
		for _, wsPt := range webShop.PaymentTypes{
			retMyPaymentTypes = append(retMyPaymentTypes, wsPt.Name)
		}
	}

	ret.PaymentOptions = retPaymentTypes
	ret.MyPaymentOptions = retMyPaymentTypes

	return ret, nil
}

func (service *Service) UpdateMyPaymentTypes(profileId string, paymentTypes []string) error {
	profileIdPrimitive, err := primitive.ObjectIDFromHex(profileId)
	if err != nil{
		return err
	}
	user, err := service.PSPRepository.GetUserByID(profileIdPrimitive)
	if err != nil{
		return err
	}
	webShopIdPrimitive, err := primitive.ObjectIDFromHex(user.WebShopId)
	if err != nil{
		return err
	}
	webShop, err := service.PSPRepository.GetWebShopByID(webShopIdPrimitive)
	if err != nil{
		return err
	}
	webShop.PaymentTypes = make([]model.PaymentType, 0)

	for _, ptName := range paymentTypes{
		pt, err := service.PSPRepository.GetPaymentTypeByName(ptName)
		if err != nil{
			return err
		}
		webShop.PaymentTypes = append(webShop.PaymentTypes, *pt)
	}
	service.PSPRepository.UpdateWebShop(webShop)

	return nil
}