package service

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ubiquitous-payment/psp/dto"
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/util"
)

func (service *Service) CreateEmptyTransaction() (string, error) {
	orderId := uuid.NewString()
	err := service.PSPRepository.CreateTransaction(&model.Transaction{ID: primitive.NewObjectID(), PSPId: orderId})
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
	t.MerchantOrderID = dto.MerchantOrderId
	t.MerchantTimestamp = dto.MerchantTimestamp
	t.PaymentMode = model.GetPaymentMode(dto.PaymentMode)
	if t.PaymentMode == model.ONE_TIME {
		t.Recurring = nil
	} else {
		t.Recurring = &model.Recurring{ID: primitive.NewObjectID(), Type: model.GetRecurringType(dto.RecurringType),
			InstallmentCount: util.String2Uint(dto.RecurringTimes), DelayedInstallmentCount: 0,
		}
	}
	t.IsSubscription = dto.IsSubscription

	return "", nil
}
