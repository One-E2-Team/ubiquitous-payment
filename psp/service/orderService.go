package service

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ubiquitous-payment/psp/dto"
	"ubiquitous-payment/psp/model"
)

func (service *Service) CreateEmptyTransaction() (string, error) {
	orderId := uuid.NewString()
	err := service.PSPRepository.CreateTransaction(&model.Transaction{ID:primitive.NewObjectID(),PSPId: orderId})
	if err != nil {
		return "", err
	}
	return orderId, nil
}

func (service *Service) FillTransaction(dto dto.WebShopOrderDTO) (string, error) {
	t, err := service.PSPRepository.GetTransactionByPspId(dto.PspOrderId)
	if err != nil {
		return "", err
	}
	//TODO: WebshopId from token
	t.Amount = dto.Amount
	t.Currency = dto.Currency
	t.SuccessURL = dto.SuccessUrl
	t.FailURL = dto.FailedUrl
	t.ErrorURL = dto.ErrorUrl
	t.PaymentMode =	model.GetPaymentModeType(dto.PaymentMode)
	return "", nil
}

