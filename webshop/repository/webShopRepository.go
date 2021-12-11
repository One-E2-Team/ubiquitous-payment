package repository

import "ubiquitous-payment/webshop/model"

func (repo *Repository) GetFirstWebShop() (*model.WebShop, error) {
	webShop := &model.WebShop{}
	if err := repo.RelationalDatabase.Table("web_shops").First(&webShop).Error; err != nil {
		return nil, err
	}
	return webShop, nil
}

func (repo *Repository) UpdateWebShop(webShop *model.WebShop) error {
	return repo.RelationalDatabase.Save(webShop).Error
}
