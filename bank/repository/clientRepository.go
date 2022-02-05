package repository

import (
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

func (repo *Repository) GetPanNumbersByClientId(clientId uint) ([]string, error) {
	panNumbers := make([]string, 0)
	if err := repo.Database.Table("client_accounts").Raw("select cc.pan from credit_cards cc where cc.id in"+
		"(select ac.credit_card_id from account_cards ac where ac.client_account_id in"+
		"(select ua.client_account_id from user_accounts ua where ua.client_id = ?))", clientId).Scan(&panNumbers).Error; err != nil {
		return nil, err
	}
	return panNumbers, nil
}
