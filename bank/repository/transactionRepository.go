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

func (repo *Repository) GetTransactionByPaymentUrlId(paymentUrlId string) (*model.Transaction, error) {
	transaction := &model.Transaction{}
	if err := repo.Database.First(&transaction, "payment_url_id = ?", paymentUrlId).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (repo *Repository) GetTransactionByPaymentId(paymentId string) (*model.Transaction, error) {
	transaction := &model.Transaction{}
	if err := repo.Database.First(&transaction, "payment_id = ?", paymentId).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (repo *Repository) GetClientTransactions(accountNumber string, panNumbers []string) ([]model.Transaction, error) {
	transactions := make([]model.Transaction, 0)
	if err := repo.Database.Raw("select * from transactions where merchant_id = ? "+
		"or issuer_pan in (?)", accountNumber, panNumbers).Scan(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}
