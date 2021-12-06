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
	orderId, err := service.getOrderIdFrom()
	fmt.Println(orderId)
	// TODO: create webShopOrder
	order := &model.Order{Timestamp: time.Now(), BuyerProfileId: loggedUserId, ProductId: productID}
	return service.WSRepository.CreateOrder(order)
}

func (service *Service) getOrderIdFrom() (string, error) {
	pspHost, pspPort := util.GetPSPHostAndPort()
	resp, err := util.PSPRequest(http.MethodGet,
		util.GetPSPProtocol()+"://"+pspHost+":"+pspPort+"/api/psp/order-id",
		nil, map[string]string{})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(resp)
	// TODO: extract id from response
	return "Good", nil
}
