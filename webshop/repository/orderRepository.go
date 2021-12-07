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

