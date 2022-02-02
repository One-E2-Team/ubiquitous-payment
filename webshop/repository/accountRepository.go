package repository

import (
	"errors"
	"fmt"
	"ubiquitous-payment/webshop/model"
)

func (repo *Repository) GetAccountsByProfileId(profileId uint) ([]model.Account, error) {
	var accounts []model.Account
	result := repo.RelationalDatabase.Table("accounts").Select("*").Where("profile_id=?", profileId).Find(&accounts)
	return accounts, result.Error
}

func (repo *Repository) GetAccountsByProfileIdAndPaymentType(profileId uint, paymentTypeId uint) ([]model.Account, error) {
	var accounts []model.Account
	result := repo.RelationalDatabase.Table("accounts").Select("*").Where("profile_id=?", profileId).Where("payment_type_id = ?", paymentTypeId).Find(&accounts)
	return accounts, result.Error
}

func (repo *Repository) CreateAccount(account *model.Account) error {
	result := repo.RelationalDatabase.Create(account)
	if result.RowsAffected == 0 {
		return errors.New("account was not created")
	}
	return nil
}

func (repo *Repository) DeleteAccount(id uint) error {
	var account model.Account
	result := repo.RelationalDatabase.Find(&account,"id = ?", id)
	if result.Error != nil{
		return nil
	}
	fmt.Println(account)
	res := repo.RelationalDatabase.Delete(&account)
	if res.Error != nil{
		return nil
	}
	if res.RowsAffected == 0 {
		return errors.New("account was not deleted")
	}
	return nil
}

func (repo *Repository) GetAccountById(id uint) *model.Account {
	var ret model.Account
	result := repo.RelationalDatabase.Table("accounts").First(&ret, "id = ?", id)
	if result.Error != nil{
		return nil
	}
	return &ret
}