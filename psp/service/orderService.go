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

func (service *Service) FillTransaction(dto dto.WebShopOrderDTO, webShopName string) (string, error) {
	t, err := service.PSPRepository.GetTransactionByPspId(dto.PspOrderId)
	if err != nil {
		return "", err
	}
	t.WebShopID = webShopName
	t.Amount = dto.Amount
	t.Currency = dto.Currency
	t.SuccessURL = dto.SuccessUrl
	t.FailURL = dto.FailUrl
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
	t.MerchantAccounts, t.AvailablePaymentTypes, err = service.extractAccounts(dto.PaymentTo)
	if err != nil {
		return "", err
	}
	err = service.PSPRepository.UpdateTransaction(t)
	//TODO: redirect link for choosing PSP payment
	return "", err
}

func (service *Service) SelectPaymentType(request dto.SelectedPaymentTypeDTO) error {
	t, err := service.PSPRepository.GetTransactionById(request.ID)
	if err != nil {
		return err
	}
	pt, err := service.PSPRepository.GetPaymentTypeByName(request.PaymentTypeName)
	if err != nil {
		return nil
	}
	t.SelectedPaymentType = *pt
	err = service.PSPRepository.UpdateTransaction(t)
	return err
}

func (service *Service) GetAvailablePaymentTypeNames(transactionID string) ([]string, error) {
	paymentTypes, err := service.PSPRepository.GetAvailablePaymentTypes(transactionID)
	if err != nil {
		return nil, err
	}
	var paymentTypeNames []string
	for _, paymentType := range paymentTypes {
		paymentTypeNames = append(paymentTypeNames, paymentType.Name)
	}
	return paymentTypeNames, nil
}

func (service *Service) extractAccounts(paymentData map[string][]string) ([]model.Account, []model.PaymentType, error) {
	accounts := make([]model.Account, 0)
	avPaymentTypes := make([]model.PaymentType, 0)
	allPaymentTypes, err := service.PSPRepository.GetAllPaymentTypes()
	if err != nil {
		return nil, nil, err
	}
	for name, accData := range paymentData {
		for _, pt := range allPaymentTypes {
			if name == pt.Name {
				acc := model.Account{ID: primitive.NewObjectID(), AccountID: accData[0],
					Secret: accData[1], PaymentType: pt}
				accounts = append(accounts, acc)
				if !paymentTypeListContains(pt, avPaymentTypes) {
					avPaymentTypes = append(avPaymentTypes, pt)
				}
			}
		}
	}
	return accounts, avPaymentTypes, err
}

func paymentTypeListContains(paymentType model.PaymentType, list []model.PaymentType) bool {
	for _, pt := range list {
		if pt.Name == paymentType.Name {
			return true
		}
	}
	return false
}
