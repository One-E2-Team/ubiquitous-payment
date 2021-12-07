package service

import (
	"fmt"
	"net/http"
	"time"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/model"
)

func (service *Service) CreateOrder(productID uint, loggedUserId uint) error {
	_, err := service.WSRepository.GetProduct(productID)
	if err != nil {
		return err
	}
	pspId, err := service.getOrderIdFromPSP()
	if err != nil {
		return err
	}
	order := &model.Order{Timestamp: time.Now(), BuyerProfileId: loggedUserId, ProductId: productID}
	err = service.WSRepository.CreateOrder(order)
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
