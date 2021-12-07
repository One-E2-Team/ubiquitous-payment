package repository

import "ubiquitous-payment/webshop/model"

func (repo *Repository) GetAccountsByProfileId(profileId uint) ([]model.Account, error) {
	var accounts []model.Account
	result := repo.RelationalDatabase.Table("accounts").Select("*").Where("profile_id=?", profileId).Find(&accounts)
	return accounts, result.Error
}
