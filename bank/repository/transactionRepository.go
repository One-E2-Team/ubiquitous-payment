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

func (repo *Repository) GetAcquirerTransactions(accountNumber string) ([]model.Transaction, error) {
	transactions := make([]model.Transaction, 0)
	if err := repo.Database.Raw("select * from transactions where merchant_id = ? ", accountNumber).Scan(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (repo *Repository) GetIssuerTransactions(panNumbers []string) ([]model.Transaction, error) {
	transactions := make([]model.Transaction, 0)
	if err := repo.Database.Raw("select * from transactions where issuer_pan in (?)", panNumbers).Scan(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (repo *Repository) GetAllTransactions() ([]model.Transaction, error) {
	transactions := make([]model.Transaction, 0)
	if err := repo.Database.Raw("select * from transactions").Scan(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}
