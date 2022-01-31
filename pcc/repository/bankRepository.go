package repository

import "ubiquitous-payment/pcc/model"

func (repo *Repository) GetBankByPanPrefix(panPrefix string) (*model.Bank, error){
	bank := &model.Bank{}
	err := repo.RelationalDatabase.First(&bank, "pan_prefix = ?", panPrefix)
	return bank, err.Error
}
