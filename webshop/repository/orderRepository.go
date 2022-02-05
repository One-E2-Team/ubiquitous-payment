package repository

import (
	"errors"
	"ubiquitous-payment/webshop/model"
)

func (repo *Repository) CreateOrder(order *model.Order) error {
	result := repo.RelationalDatabase.Create(order)
	if result.RowsAffected == 0 {
		return errors.New("order was not created")
	}
	return nil
}

func (repo *Repository) CreatePspOrder(pspOrder *model.PSPOrder) error {
	result := repo.RelationalDatabase.Create(pspOrder)
	if result.RowsAffected == 0 {
		return errors.New("psp order was not created")
	}
	return nil
}

func (repo *Repository) GetMyOrders(profileId uint) ([]model.Order ,error) {
	var myOrders []model.Order
	result := repo.RelationalDatabase.Find(&myOrders, "buyer_profile_id = ?", profileId)
	if result.Error != nil {
		return nil, errors.New("error in query for getting my orders")
	}
	return myOrders,nil
}

func (repo *Repository) GetPspOrderByOrderId(orderId uint) (model.PSPOrder ,error) {
	var pspOrder model.PSPOrder
	result := repo.RelationalDatabase.Table("psp_orders").Last(&pspOrder, "order_id = ?", orderId)
	return pspOrder, result.Error
}

func (repo *Repository) GetPspOrderByPspId(pspId string) (model.PSPOrder ,error) {
	var pspOrder model.PSPOrder
	result := repo.RelationalDatabase.Table("psp_orders").Last(&pspOrder, "psp_id = ?", pspId)
	return pspOrder, result.Error
}

func (repo *Repository) GetOrdersByProductId(productId uint) ([]model.Order ,error) {
	var orders []model.Order
	result := repo.RelationalDatabase.Find(&orders, "product_id = ?", productId)
	if result.Error != nil {
		return nil, errors.New("error in query for getting my orders")
	}
	return orders,nil
}
