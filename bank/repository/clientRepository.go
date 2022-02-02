package repository

import (
	"ubiquitous-payment/bank/model"
)

func (repo *Repository) GetClientAccount(accountNumber string) (*model.ClientAccount, error) {
	clientAccount := &model.ClientAccount{}
	if err := repo.Database.First(&clientAccount, "account_number = ?", accountNumber).Error; err != nil {
		return nil, err
	}
	return clientAccount, nil
}

func (repo *Repository) GetCreditCard(pan string) (*model.CreditCard, error) {
	creditCard := &model.CreditCard{}
	if err := repo.Database.First(&creditCard, "pan = ?", pan).Error; err != nil {
		return nil, err
	}
	return creditCard, nil
}

func (repo *Repository) GetClientAccountByPan(pan string) (*model.ClientAccount, error) {
	clientAccount := &model.ClientAccount{}
	if err := repo.Database.Table("client_accounts").Raw("select * from client_accounts ca where ca.id ="+
		"(select ac.client_account_id from account_cards ac "+
		"where ac.credit_card_id = "+
		"(select cc.id from credit_cards cc where cc.pan = ?))", pan).Scan(&clientAccount).Error; err != nil {
		return nil, err
	}
	return clientAccount, nil
}

func (repo *Repository) GetRoleByName(name string) (*model.Role, error) {
	role := &model.Role{}
	if err := repo.Database.Table("roles").First(&role, "name = ?", name).Error; err != nil {
		return nil, err
	}
	return role, nil
}
