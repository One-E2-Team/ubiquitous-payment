package service

import (
	"time"
	"ubiquitous-payment/webshop/model"
)

func (service *Service) CreateOrder(productID uint, loggedUserId uint) error {
	order := &model.Order{Timestamp: time.Now(), BuyerProfileId: loggedUserId, ProductId: productID}
	return service.WSRepository.CreateOrder(order)
}
