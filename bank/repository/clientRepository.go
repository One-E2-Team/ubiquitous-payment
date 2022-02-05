package repository

import (
	"errors"
	"gorm.io/gorm/clause"
	"ubiquitous-payment/bank/model"
)

func (repo *Repository) GetPrivilegesByClientId(clientId uint) (*[]string, error) {
	var privileges []string
	if err := repo.Database.Raw("select p.name from privileges p, role_privileges rp "+
		"where rp.role_id in (select r.id from roles r, user_roles ur "+
		"where ur.client_id = ? and ur.role_id = r.id)"+
		"and p.id = rp.privilege_id", clientId).Scan(&privileges).Error; err != nil {
		return nil, err
	}
	return &privileges, nil
}

func (repo *Repository) GetClientByUsername(username string) (*model.Client, error) {
	client := &model.Client{}
	if err := repo.Database.Preload("Roles").Table("clients").First(&client, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return client, nil
}

func (repo *Repository) GetClientById(clientId uint) (*model.Client, error) {
	client := &model.Client{}
	result := repo.Database.Preload("Accounts.CreditCards").Preload(clause.Associations).First(&client, "id = ?", clientId)
	return client, result.Error
}

func (repo *Repository) GetClientAccount(accountNumber string) (*model.ClientAccount, error) {
	clientAccounts := make([]model.ClientAccount, 0)
	if err := repo.Database.Raw("select * from client_accounts").Scan(&clientAccounts).Error; err != nil {
		return nil, err
	}

	for _, clientAccount := range clientAccounts {
		if clientAccount.AccountNumber.Data == accountNumber {
			return &clientAccount, nil
		}
	}
	return nil, errors.New("no client account found")
}

func (repo *Repository) GetCreditCard(pan string) (*model.CreditCard, error) {
	creditCards := make([]model.CreditCard, 0)
	if err := repo.Database.Raw("select * from credit_cards").Scan(&creditCards).Error; err != nil {
		return nil, err
	}

	for _, creditCard := range creditCards {
		if creditCard.Pan.Data == pan {
			return &creditCard, nil
		}
	}
	return nil, errors.New("no credit card found")
}

func (repo *Repository) GetClientAccountByPan(pan string) (*model.ClientAccount, error) {
	clientAccount := &model.ClientAccount{}

	creditCard, err := repo.GetCreditCard(pan)
	if err != nil {
		return nil, err
	}

	if err := repo.Database.Table("client_accounts").Raw("select * from client_accounts ca where ca.id ="+
		"(select ac.client_account_id from account_cards ac "+
		"where ac.credit_card_id = ?)", creditCard.ID).Scan(&clientAccount).Error; err != nil {
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

func (repo *Repository) GetPanNumbersByClientId(clientId uint) ([]string, error) {
	client, err := repo.GetClientById(clientId)
	if err != nil {
		return nil, err
	}

	panNumbers := make([]string, 0)
	for _, account := range client.Accounts {
		for _, creditCard := range account.CreditCards {
			panNumbers = append(panNumbers, creditCard.Pan.Data)
		}
	}

	return panNumbers, nil
}
