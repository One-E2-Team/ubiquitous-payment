package service

import (
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ubiquitous-payment/psp-plugins/pspdto"
	"ubiquitous-payment/psp-plugins/pspdto/mapper"
	"ubiquitous-payment/psp/dto"
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/psputil"
	"ubiquitous-payment/util"
)

func (service *Service) CreateEmptyTransaction() (string, error) {
	orderId := uuid.NewString()
	err := service.PSPRepository.Create(&model.Transaction{ID: primitive.NewObjectID(), PSPId: orderId}, psputil.TransactionsCollectionName)
	if err != nil {
		return "", err
	}
	return orderId, nil
}

func (service *Service) FillTransaction(dto dto.WebShopOrderDTO, webShopOwnerID string) (string, error) {
	t, err := service.PSPRepository.GetTransactionByPspId(dto.PspOrderId)
	if err != nil {
		return "", err
	}
	user, err := service.PSPRepository.GetUserByID(util.String2MongoID(webShopOwnerID))
	if err != nil {
		return "", err
	}
	webShop, err := service.PSPRepository.GetWebShopByID(util.String2MongoID(user.WebShopId))
	t.WebShopID = webShop.Name
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
			InstallmentCount: util.String2Uint(dto.RecurringTimes), DelayedInstallmentCount: dto.DelayedInstallments,
		}
	}
	t.IsSubscription = dto.IsSubscription
	t.MerchantAccounts, t.AvailablePaymentTypes, err = service.extractAccounts(dto.PaymentTo)
	if err != nil {
		return "", err
	}
	err = service.PSPRepository.UpdateTransaction(t)

	logContent := "Created transaction: '" + util.MongoID2String(t.ID) + "'"
	util.Logging(util.SUCCESS, "Service.FillTransaction", logContent, "psp")

	pspFrontHost, pspFrontPort := util.GetPSPFrontHostAndPort()
	return util.GetPSPProtocol() + "://" + pspFrontHost + ":" + pspFrontPort + "/#/choose-payment-type/" + t.ID.Hex(), err
}

func (service *Service) SelectPaymentType(request dto.SelectedPaymentTypeDTO) (*pspdto.TransactionCreatedDTO, error) {
	id, err := primitive.ObjectIDFromHex(request.ID)
	if err != nil {
		return nil, err
	}
	t, err := service.PSPRepository.GetTransactionById(id)
	if err != nil {
		return nil, err
	}
	pt, err := service.PSPRepository.GetPaymentTypeByName(request.PaymentTypeName)
	if err != nil {
		return nil, nil
	}
	t.SelectedPaymentType = *pt
	err = service.PSPRepository.UpdateTransaction(t)
	if err != nil {
		return nil, err
	}

	logContent := "For transaction: '" + request.ID + "' chosen payment type is '" + request.PaymentTypeName + "'"
	util.Logging(util.INFO, "Service.SelectPaymentType", logContent, "psp")

	result, err := service.ExecuteTransaction(t)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *Service) ExecuteTransaction(t *model.Transaction) (*pspdto.TransactionCreatedDTO, error) {
	plugin, err := psputil.GetPlugin(t.SelectedPaymentType.Name)
	if err != nil {
		return nil, err
	}
	transactionDTO, err := mapper.TransactionToTransactionDTO(*t, plugin)
	result, err := plugin.ExecuteTransaction(transactionDTO)
	if err != nil {
		t.TransactionStatus = model.FAILED
		_ = service.PSPRepository.UpdateTransaction(t)
		return nil, err
	}
	t.ExternalTransactionId = result.TransactionId
	err = service.PSPRepository.UpdateTransaction(t)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (service *Service) UpdateTransactionSuccess(transactionID string) (string, error) {
	return service.updateTransactionStatus(transactionID, model.FULFILLED)
}

func (service *Service) UpdateTransactionFail(transactionID string) (string, error) {
	return service.updateTransactionStatus(transactionID, model.FAILED)
}

func (service *Service) UpdateTransactionError(transactionID string) (string, error) {
	return service.updateTransactionStatus(transactionID, model.ERROR)
}

func (service *Service) updateTransactionStatus(externalId string, status model.TransactionStatus) (string, error) {
	logContent := "Transaction: '" + externalId + "' was '" + status.ToString() + "'"
	util.Logging(util.INFO, "Service.SelectPaymentType", logContent, "psp")
	fmt.Println("prije dobavljanja transakcije na osnovu external ida")
	t, err := service.PSPRepository.GetTransactionByExternalId(externalId)
	fmt.Println("got transaction for external id ", t)
	if err != nil {
		return "", err
	}
	//TODO: Remove paypal check
	if status == model.FULFILLED && t.SelectedPaymentType.Name != "paypal" {
		plugin, err := psputil.GetPlugin(t.SelectedPaymentType.Name)
		if err != nil {
			return "", err
		}
		plan := t.IsSubscription || (t.Recurring != nil)
		isFulfilled, err := plugin.CaptureTransaction(t.ExternalTransactionId, plan)
		fmt.Println("fulfilled je ", isFulfilled)
		if err != nil || !isFulfilled {
			t.TransactionStatus = model.ERROR
		}
		if isFulfilled {
			t.TransactionStatus = status
		}
	} else {
		t.TransactionStatus = status
	}
	return t.GetURLByStatus(), service.PSPRepository.UpdateTransaction(t)
}

func (service *Service) GetAvailablePaymentTypeNames(transactionID string) ([]string, error) {
	t, err := service.PSPRepository.GetTransactionById(util.String2MongoID(transactionID))
	if err != nil {
		return nil, err
	}
	paymentTypes, err := service.PSPRepository.GetAvailablePaymentTypes(transactionID)
	if err != nil {
		return nil, err
	}
	var paymentTypeNames []string
	for _, paymentType := range paymentTypes {
		if t.IsSubscription || (t.Recurring != nil) {
			plugin, err := psputil.GetPlugin(paymentType.Name)
			if err != nil {
				return nil, err
			}
			if plugin.SupportsPlanPayment() {
				paymentTypeNames = append(paymentTypeNames, paymentType.Name)
			}
		} else {
			paymentTypeNames = append(paymentTypeNames, paymentType.Name)
		}
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
				if !service.paymentTypeListContains(avPaymentTypes, pt.Name) {
					avPaymentTypes = append(avPaymentTypes, pt)
				}
			}
		}
	}
	return accounts, avPaymentTypes, err
}

func (service *Service) CheckForPaymentBitcoin(id string) (*dto.CheckForPaymentDTO, error) {
	plugin, err := psputil.GetPlugin("bitcoin")
	if err != nil {
		return nil, err
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	t, err := service.PSPRepository.GetTransactionById(objectId)
	if err != nil {
		return nil, err
	}
	plan := t.IsSubscription || (t.Recurring != nil)
	isCaptured, err := plugin.CaptureTransaction(t.ExternalTransactionId, plan)
	if err != nil {
		return nil, err
	}
	result := dto.CheckForPaymentDTO{PaymentCaptured: isCaptured, SuccessUrl: t.SuccessURL}
	return &result, nil
}
