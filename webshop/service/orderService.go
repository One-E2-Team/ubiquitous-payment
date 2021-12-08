package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/model"
)

func (service *Service) CreateOrder(productID uint, loggedUserId uint) error {
	product, err := service.WSRepository.GetProduct(productID)
	if err != nil {
		return err
	}
	pspId, err := service.getOrderIdFromPSP()
	if err != nil {
		return err
	}
	order := &model.Order{Timestamp: time.Now(), UUID: uuid.NewString(), BuyerProfileId: loggedUserId, ProductId: productID}
	err = service.WSRepository.CreateOrder(order)
	if err != nil {
		return err
	}
	_, err = service.getRedirectLinkFromPsp(product, order, pspId)
	if err != nil {
		return err
	}
	pspOrder := &model.PSPOrder{PSPId: pspId, OrderId: order.ID, Timestamp: time.Now(), OrderStatus: model.PLACED}
	err = service.WSRepository.CreatePspOrder(pspOrder)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) getOrderIdFromPSP() (string, error) {
	pspHost, pspPort := util.GetPSPHostAndPort()
	resp, err := util.PSPRequest(http.MethodGet,
		util.GetPSPProtocol()+"://"+pspHost+":"+pspPort+"/api/psp/order-id",
		nil, map[string]string{})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var orderId string
	err = util.UnmarshalResponse(resp, &orderId)
	if err != nil {
		return "", err
	}
	return orderId, nil
}

func (service *Service) getRedirectLinkFromPsp(product *model.Product, order *model.Order, pspId string) (string, error) {
	paymentData, err := service.getPaymentData(product.MerchantProfileId)
	if err != nil {
		return "", err
	}

	message := make(map[string]interface{})
	message["pspOrderId"] = pspId
	message["amount"] = product.Price
	message["currency"] = product.Currency
	message["paymentMode"] = "ONE_TIME"
	message["isSubscription"] = false
	message["recurringType"] = ""
	switch product.NumOfInstallments {
	case 0:
		message["recurringTimes"] = "0"
		message["isSubscription"] = true
	case 1:
		message["recurringTimes"] = "1"
	default:
		message["paymentMode"] = "RECURRING"
		message["recurringType"] = product.RecurringType.String()
		message["recurringTimes"] = util.Uint2String(product.NumOfInstallments)
	}
	message["paymentTo"] = paymentData
	message["successUrl"] = ""
	message["failUrl"] = ""
	message["errorUrl"] = ""
	message["merchantTimestamp"] = order.Timestamp
	message["merchantOrderId"] = order.UUID

	data, _ := json.Marshal(message)
	pspHost, pspPort := util.GetPSPHostAndPort()
	resp, err := util.PSPRequest(http.MethodPost,
		util.GetPSPProtocol()+"://"+pspHost+":"+pspPort+"/api/order",
		data, map[string]string{})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var redirectLink string
	err = util.UnmarshalResponse(resp, &redirectLink)
	if err != nil {
		return "", err
	}
	return redirectLink, nil
}

func (service *Service) getPaymentData(merchantId uint) (map[string]interface{}, error) {
	merchantAccounts, err := service.WSRepository.GetAccountsByProfileId(merchantId)
	if err != nil {
		return nil, err
	}
	validPaymentTypes, err := service.WSRepository.GetValidPaymentTypes()
	if err != nil {
		return nil, err
	}
	ret := make(map[string]interface{})
	for _, pt := range validPaymentTypes {
		for _, acc := range merchantAccounts {
			if pt.ID == acc.PaymentTypeId {
				ret[pt.Name] = [2]string{acc.AccountID, acc.Secret}
				continue
			}
		}
	}
	return ret, nil
}
