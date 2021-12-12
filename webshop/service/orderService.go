package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/model"
	"ubiquitous-payment/webshop/wsutil/pspAuth"
)

func (service *Service) CreateOrder(productID uint, loggedUserId uint) (string, error) {
	product, err := service.WSRepository.GetProduct(productID)
	if err != nil {
		return "", err
	}
	pspId, err := service.getOrderIdFromPSP()
	if err != nil {
		return "", err
	}
	order := &model.Order{Timestamp: time.Now(), UUID: uuid.NewString(), BuyerProfileId: loggedUserId, ProductId: productID}
	err = service.WSRepository.CreateOrder(order)
	if err != nil {
		return "", err
	}

	logContent := "User: '" + util.Uint2String(loggedUserId) + "' created order: '" + util.Uint2String(order.ID) + "' for product: '" + util.Uint2String(productID) + "'"
	util.Logging(util.INFO, "Service.CreateOrder", logContent, "web-shop")

	redirectUrl, err := service.getRedirectLinkFromPsp(product, order, pspId)
	if err != nil {
		return "", err
	}
	pspOrder := &model.PSPOrder{PSPId: pspId, OrderId: order.ID, Timestamp: time.Now(), OrderStatus: model.PLACED}
	err = service.WSRepository.CreatePspOrder(pspOrder)
	if err != nil {
		return "", err
	}

	logContent = "User: '" + util.Uint2String(loggedUserId) + "' created PSP order" + util.Uint2String(pspOrder.ID) + "' for order: '" + util.Uint2String(order.ID) + "'"
	util.Logging(util.SUCCESS, "Service.CreateOrder", logContent, "web-shop")
	
	return redirectUrl, nil
}

func (service *Service) getOrderIdFromPSP() (string, error) {
	resp, err := pspAuth.PSPRequest(http.MethodGet, "/api/psp/order-id",
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
	message["delayedInstallments"] = product.DelayedInstallments
	switch product.NumOfInstallments {
	case 0:
		message["recurringTimes"] = "0"
		message["isSubscription"] = true
	case 1:
		message["recurringTimes"] = "1"
	default:
		message["paymentMode"] = "RECURRING"
		message["recurringType"] = product.RecurringType
		message["recurringTimes"] = util.Uint2String(product.NumOfInstallments)
	}
	message["paymentTo"] = paymentData
	wsFrontHost, wsFrontPort := util.GetWebShopFrontHostAndPort()
	initialUrl := util.GetWebShopProtocol() + "://" + wsFrontHost + ":" + wsFrontPort + "/#/order"
	message["successUrl"] = initialUrl + "/success/" + pspId
	message["failUrl"] = initialUrl + "/fail" + pspId
	message["errorUrl"] = initialUrl + "/error" + pspId
	message["merchantTimestamp"] = order.Timestamp
	message["merchantOrderId"] = order.UUID

	data, _ := json.Marshal(message)
	resp, err := pspAuth.PSPRequest(http.MethodPost, "/api/order",
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
