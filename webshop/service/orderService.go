package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/dto"
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

func (service *Service) GetMyOrders(profileId uint) ([]dto.MyOrderDTO, error) {
	retOrders := make([]dto.MyOrderDTO, 0)
	myOrders, err := service.WSRepository.GetMyOrders(profileId)
	if err != nil {
		return nil, err
	}
	for _, order := range myOrders {
		newOrderDto := dto.MyOrderDTO{OrderId: order.ID, Timestamp: order.Timestamp}

		product, err := service.WSRepository.GetProduct(order.ProductId)
		if err != nil {
			return nil, err
		}
		newOrderDto.ProductName = product.Name
		newOrderDto.ProductPrice = product.Price
		newOrderDto.Currency = product.Currency
		newOrderDto.NumberOfInstallments = product.NumOfInstallments
		newOrderDto.NumberOfInstallments = product.DelayedInstallments
		newOrderDto.RecurringType = string(product.RecurringType)

		paymentType, err := service.WSRepository.GetPaymentTypeById(order.PaymentTypeId)
		if err != nil {
			return nil, err
		}
		newOrderDto.PaymentType = paymentType.Name

		pspOrder, err := service.WSRepository.GetPspOrderByOrderId(order.ID)
		newOrderDto.PSPId = pspOrder.PSPId
		newOrderDto.OrderStatus = string(pspOrder.OrderStatus)

		retOrders = append(retOrders, newOrderDto)
	}

	return retOrders, nil
}

func (service *Service) UpdatePspOrder(pspId string, status string) error {
	pspOrder, err := service.WSRepository.GetPspOrderByPspId(pspId)
	if err != nil {
		return err
	}
	pspOrder.OrderStatus = model.OrderStatus(status)
	return service.WSRepository.Update(pspOrder)
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
	message["failUrl"] = initialUrl + "/fail/" + pspId
	message["errorUrl"] = initialUrl + "/error/" + pspId
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
				ret[pt.Name] = [2]string{acc.AccountID.Data, acc.Secret.Data}
				continue
			}
		}
	}
	return ret, nil
}

func (service *Service) GetSellersOrders(id uint) ([]dto.MyOrderDTO, error) {
	ordersDto := make([]dto.MyOrderDTO, 0)

	myProducts, err := service.WSRepository.GetProductsByMerchantId(id)
	if err != nil {
		return nil, err
	}
	fmt.Println(myProducts)

	for _, prod := range myProducts {
		orders, err := service.WSRepository.GetOrdersByProductId(prod.ID)
		if err != nil {
			return nil, err
		}
		for _, order := range orders {
			pspOrder, err := service.WSRepository.GetPspOrderByOrderId(order.ID)
			if err != nil {
				return nil, err
			}
			paymentType, err := service.WSRepository.GetPaymentTypeById(order.PaymentTypeId)
			if err != nil {
				return nil, err
			}
			orderDto := dto.MyOrderDTO{OrderId: order.ID, Timestamp: order.Timestamp, ProductName: prod.Name,
				ProductPrice: prod.Price, Currency: prod.Currency, PaymentType: paymentType.Name,
				PSPId: pspOrder.PSPId, OrderStatus: string(pspOrder.OrderStatus),
				NumberOfInstallments: prod.NumOfInstallments, DelayedInstallments: prod.DelayedInstallments,
				RecurringType: string(prod.RecurringType)}
			ordersDto = append(ordersDto, orderDto)
		}
	}

	return ordersDto, nil
}
