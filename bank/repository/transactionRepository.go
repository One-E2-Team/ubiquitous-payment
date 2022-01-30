package repository

import (
	"fmt"
	"ubiquitous-payment/bank/model"
)

func (repo *Repository) CreateTransaction(transaction *model.Transaction) error {
	result := repo.Database.Create(transaction)
	if result.RowsAffected == 0 {
		return fmt.Errorf("transaction not created")
	}
	return nil
}
