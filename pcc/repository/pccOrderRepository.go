package repository

import (
	"errors"
	"ubiquitous-payment/pcc/model"
)

func (repo *Repository) CreatePccOrder(pccOrder *model.PccOrder) error{
	result := repo.RelationalDatabase.Create(pccOrder)
	if result.RowsAffected == 0 {
		return errors.New("order was not created")
	}
	return nil
}
