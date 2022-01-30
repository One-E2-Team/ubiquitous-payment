package repository

import "ubiquitous-payment/bank/model"

func (repo *Repository) GetClientAccount(accountNumber string) (*model.ClientAccount, error) {
	clientAccount := &model.ClientAccount{}
	if err := repo.Database.First(&clientAccount, "account_number = ?", accountNumber).Error; err != nil {
		return nil, err
	}
	return clientAccount, nil
}
